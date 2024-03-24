package utils

import (
	"bingotest01/application/constants"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/pkg/sftp"
	"go.uber.org/zap"
	"golang.org/x/crypto/ssh"
	"strings"
	"time"
)

//ssh客户端连接信息的结构体

type SSH struct {
	IP         string       //IP地址
	Port       int          //端口号
	Username   string       //用户名
	Mod        string       //认证方式[password:密码,key:秘钥认证]
	Password   string       //密码
	Key        string       //认证私钥
	Client     *ssh.Client  //ssh客户端
	SftpClient *sftp.Client //sftp客户端
	//console
	Session    *ssh.Session //ssh回话对象
	Channel    ssh.Channel  //ssh通信信道
	LastResult string       //最近一次执行命令的结果
}

func NewSSH(ip, username, password, key, mod string, port ...int) *SSH {
	/*
		创建命令行实例
		@param ip IP地址
		@param username 用户名
		@param password 登录密码
		@param key 认证私钥
		@param mode 认证模式(password:密码 | key:秘钥)
		@param port 端口号,不填写默认为22
	*/

	//实例化client对象
	client := new(SSH)
	client.IP = ip
	client.Username = username
	client.Password = password
	client.Key = key
	client.Mod = mod
	//判断端口
	if len(port) <= 0 {
		client.Port = 22
	} else {
		client.Port = port[0]
	}
	return client
}

//ssh连接

func (s *SSH) Connect() error {

	//ssh连接
	config := ssh.ClientConfig{
		User:            s.Username,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         10 * time.Second,
	}

	//判断SSH连接的认证方式
	if s.Mod == "key" {
		signer, err := ssh.ParsePrivateKey([]byte(s.Key))
		if err != nil {
			zap.S().Fatalf("ssh key signer faild: %v", err)
		}
		config.Auth = []ssh.AuthMethod{ssh.PublicKeys(signer)}
	} else {
		//fmt.Println("s.password: ", s.Password)
		config.Auth = []ssh.AuthMethod{ssh.Password(s.Password)}
	}

	//创建ssh客户端
	addr := fmt.Sprintf("%s:%d", s.IP, s.Port)
	sshClient, err := ssh.Dial("tcp", addr, &config)
	if err != nil {
		return err
	}
	s.Client = sshClient

	//创建一个客户端SSH回话
	session, err := s.Client.NewSession()
	if err != nil {
		return err
	}
	s.Session = session
	return nil
}

//执行ssh命令

func (s SSH) Run(command string) (string, error) {
	//执行ssh命令
	//@param command 要执行的命令,多个命令采用;分隔
	if s.Client == nil {
		if err := s.Connect(); err != nil {
			return "", err
		}
	}
	////创建一个ssh客户端
	//session, err := s.Client.NewSession()
	//if err != nil {
	//	return "", err
	//}
	defer s.Session.Close()

	//执行shell命令
	buf, err := s.Session.CombinedOutput(command)
	//记录最后一次执行命令的结果
	s.LastResult = string(buf)
	return s.LastResult, err
}

//向远程服务器添加公钥

func (s SSH) AddPublicKeyToRemoteHost(publicKey string) error {
	/*
		将公钥写入目标主机
		700 是文档拥有可读可写执行权限,同一组用户或者其他用户都不具有操作权限
		600 是文件拥有可读可写,不可执行,同一组用户或者其他用户都不具有操作权限
	*/
	command := fmt.Sprintf("mkdir -p -m 700 ~/.ssh && echo %v >> ~/.ssh/authorized_keys && chmod 600 ~/.ssh/authorized_keys", strings.TrimSpace(publicKey))
	//调用执行命令的方法
	_, err := s.Run(command)
	if err != nil {
		message := fmt.Sprintf("%v: %v", constants.AddPublicKeyFail, err)
		return errors.New(message)
	}
	return nil
}

//websocket

type MyReader struct {
}

var cmdChannel = make(chan string)

func (r MyReader) Read(p []byte) (n int, err error) {
	var cmd string
	cmd = <-cmdChannel
	cmdB := []byte(cmd + "\n")
	for i, v := range cmdB {
		p[i] = v
	}
	n = len(cmdB)
	return n, err
}

type MyWriter struct {
}

var resChannel = make(chan string)

func (w MyWriter) Write(p []byte) (n int, err error) {
	res := string(p)
	resChannel <- res
	return len(p), err
}

func (s *SSH) WS2SSH(ws *websocket.Conn) {

	//实现websocket通信与ssh通信之间的数据转换读写
	defer func(s *SSH) {
		_ = s.Client.Close()
		_ = s.Session.Close()
	}(s)

	//重置SSH通信管道对象
	cmdChannel = make(chan string)
	resChannel = make(chan string)

	//ssh会话提供的标准输出
	s.Session.Stdout = new(MyWriter)
	//ssh会话提供的标准错误输出
	s.Session.Stderr = new(MyWriter)
	//ssh会话提供的标准输入
	s.Session.Stdin = new(MyReader)

	//配置和创建一个伪PTY终端
	fmt.Println("配置和创建一个伪pty终端")
	modes := ssh.TerminalModes{
		ssh.ECHO:          0,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	err := s.Session.RequestPty("xterm", 25, 80, modes)
	if err != nil {
		fmt.Println("Session.RequestPty err : ", err)
	}

	err = s.Session.Shell()
	if err != nil {
		fmt.Println("Session.Shell err: ", err)
	}

	//从ssh通信管道resChannel读取远程主机返回数据转发写入websocket通信管道中
	fmt.Println("启动协程1")
	go func() {
		for {
			cmd_result := <-resChannel
			fmt.Println("cmd_result>>> ", cmd_result)
			if len(cmd_result) != 0 {
				err := ws.WriteMessage(websocket.TextMessage, []byte(cmd_result))
				cmd_result = ""
				if err != nil {
					return
				}
			}
		}

	}()

	//从websocket通信管道中读取客户端发送的命令,转发给ssh通信管道
	fmt.Println("启动协程2")

	go func() {
		for {
			_, cmd, err := ws.ReadMessage()
			fmt.Println("receive cmd>>> ", string(cmd))
			if err != nil {
				return
			}
			arg := ""
			var str string
			_, _ = fmt.Sscan(string(cmd), &str, &arg)
			if len(arg) > 0 {
				arg = " " + arg
			}
			cmdChannel <- str + arg
		}
	}()

	err = s.Session.Wait()
	if err != nil {
		return
	}
}
