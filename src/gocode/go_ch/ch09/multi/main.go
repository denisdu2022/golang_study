package main

import "fmt"

type MyWriter interface {
	Write(string) error
}

type MyCloser interface {
	Close() error
}

type writerCloser struct {
	MyWriter //interface 也是一个类型,想放入一个写文件的实现,想放一个写数据库的实现
}

// 写文件
type fileWriter struct {
	filePath string
}

// 写数据库
type databaseWriter struct {
	dbName string
	dbPort string
	db     string
}

func (fw *fileWriter) Write(string) error {

	fmt.Println("write string to file")
	return nil
}

func (dw *databaseWriter) Write(string) error {
	fmt.Println("write string to file database")
	return nil
}

func (wc *writerCloser) Close() error {
	fmt.Println("close")
	return nil
}

func main() {
	//实例化写文件
	var mw MyWriter = &writerCloser{
		&databaseWriter{},
	}
	mw.Write("file")

}
