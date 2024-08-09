package main

import "fmt"

type MyWriter interface {
	Write(string)
}

type MyReader interface {
	Read() error
}

type MyReadWriter interface {
	MyWriter
	MyReader
	//也可以加自己的方法
	ReadWrite()
}

type SreadWriter struct {
}

func (s *SreadWriter) Write(s2 string) {
	//TODO implement me
	fmt.Println("write")
}

func (s *SreadWriter) Read() error {
	//TODO implement me
	fmt.Println("read")
	return nil
}

func (s *SreadWriter) ReadWrite() {
	//TODO implement me
	fmt.Println("ReadWrite")
}

func main() {
	var mrw MyReadWriter = &SreadWriter{}
	mrw.Read()
}
