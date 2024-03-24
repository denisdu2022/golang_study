package validator

import (
	"bingotest01/application/config"
	"bingotest01/application/constants"
	"bingotest01/application/model"
	"bingotest01/application/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"io/ioutil"
	"mime/multipart"
	"strconv"
	"strings"
)

//添加主机类别验证器

func HostCategoryValidator(hostCategory *model.HostCategory) error {

	//获取验证器对象
	validate, trans := utils.GenValidate()
	//验证主机类型结构体
	err := validate.Struct(hostCategory)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return errors.New(err.Translate(trans))
		}
	}
	return err
}

//主机验证器

func HostValidator(host *model.Host, ctx *gin.Context) error {

	//1.接收多文件上传的表单数据
	form, err := ctx.MultipartForm()
	if err != nil {
		return err
	}
	fmt.Println("form>>> ", form)

	//手动接收表单信息并赋值
	//form.Value接收文本字段,form.File接收上传文件

	//接收主机类别Id并赋值
	hostCategoryId, _ := strconv.Atoi(form.Value["host_category_id"][0])
	host.HostCategoryID = uint(hostCategoryId)

	//接收主机端口并赋值
	port, _ := strconv.Atoi(form.Value["port"][0])
	host.Port = uint(port)

	//接收主机用户名并赋值
	host.Username = form.Value["username"][0]
	//接收主机密码并赋值
	host.Password = form.Value["password"][0]
	//接收主机名并赋值
	host.Name = form.Value["name"][0]
	//接收主机IP地址并赋值
	host.IpAddr = form.Value["ip_addr"][0]
	//接收主机备注并赋值
	host.Remark = form.Value["remark"][0]
	fmt.Println("host>>> ", host)

	//获取验证器对象
	//基本字段的校验
	validate, trans := utils.GenValidate()
	err = validate.Struct(host)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return errors.New(err.Translate(trans))
		}
	}

	//主机类别的校验
	//判断hostCategoryID
	if host.HostCategoryID < 1 {
		return errors.New(constants.HostCategoryNotExist)
	}

	var hostCategory model.HostCategory
	err = hostCategory.GetOneById(host.HostCategoryID)
	if err != nil {
		return errors.New(constants.HostCategoryNotExist)
	}

	//验证主机是否可以ping通,可以的话,将公钥上传到服务器指定位置,方便后边的免密登录
	err = ping(host, form)
	if err != nil {
		return err
	}
	return nil
}

func ping(host *model.Host, form *multipart.Form) error {

	var err error

	//2.从表单中提取上传文件
	keys := form.File["files[]"] //务必予客户端上传时设置的字段名一致,否则接收不了
	//3.根据上传文件个数判断
	if len(keys) == 1 {
		//报错,只上传了一个文件
		return errors.New(constants.MissingKeys)
	} else if len(keys) == 2 {
		//上传密钥对两个文件
		//获取第一个文件
		fileHandlel, err := keys[0].Open()
		if err != nil {
			return err
		}
		defer fileHandlel.Close()
		//读取文件
		fileByte, _ := ioutil.ReadAll(fileHandlel)
		key1 := string(fileByte)

		//获取第二个文件
		fileHandlel2, err := keys[1].Open()
		if err != nil {
			return err
		}
		defer fileHandlel2.Close()
		fileByte2, _ := ioutil.ReadAll(fileHandlel2)
		key2 := string(fileByte2)

		//通过判断秘钥中的内容,区分公钥和私钥
		if strings.Contains(key1, "PRIVATE") {
			host.PrivateKey, host.PublicKey = key1, key2
		} else {
			host.PrivateKey, host.PublicKey = key2, key1
		}
	} else {
		//采用全局公私钥作为默认值
		host.PublicKey, host.PrivateKey = config.Conf.PublicKey, config.Conf.PrivateKey
	}

	//确认好公私钥后,创建ssh对象
	//获取ssh客户端连接对象
	cli := utils.NewSSH(host.IpAddr, host.Username, host.Password, "", "password", int(host.Port))
	//fmt.Println("cli>>> ", cli)
	//基于密码的首次连接
	if err := cli.Connect(); err != nil {
		fmt.Println("1111111111111111111111111")
		return errors.New(constants.HostInfoError)
	}

	//验证主机信息成功后,把公钥写入到远程主机
	if err := cli.AddPublicKeyToRemoteHost(host.PublicKey); err != nil {
		return errors.New(constants.SSHKeyIsInvalid)
	}
	return err
}
