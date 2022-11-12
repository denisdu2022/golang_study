package main

////按字节读取
//func readByBtes(file *os.File) {
//	data := make([]byte, 9)
//	file.Read(data)
//	fmt.Println(data)
//	fmt.Println(string(data))
//}
//
////按行读
//func readByLine(file *os.File) {
//	reader := bufio.NewReader(file)
//	//lineContent, _ := reader.ReadString('\n') //以换行符结束
//	//fmt.Println(lineContent)
//	lineContent, _, _ := reader.ReadLine() //按行读
//	fmt.Println(lineContent)
//	fmt.Println(string(lineContent))
//}
//
////按行读,全部读出
//func readByLine(file *os.File) {
//	reader := bufio.NewReader(file)
//	for true {
//		lineContent, err := reader.ReadString('\n') //以换行符结束
//		fmt.Print(lineContent)
//		if err == io.EOF {
//			break
//		}
//	}
//
//}
//
////读取整个文件
//func readByFile() {
//	content, _ := ioutil.ReadFile("渔家傲")
//	fmt.Println(content)
//	fmt.Println(string(content))
//}

func main() {
	////1. 打开文件
	//file, err := os.Open("渔家傲")
	//if err != nil {
	//	fmt.Println("err,文件打开失败: ", err)
	//} else {
	//	fmt.Println("文件打开成功...")
	//	fmt.Println("fiel: ", file)
	//}
	//
	////2. 按字节读取数据
	//data := make([]byte, 9)
	//file.Read(data)
	//fmt.Println(data)
	//fmt.Println(string(data))
	////调用按字节读取函数
	//readByBtes(file)
	//
	////2. 按行读
	//reader := bufio.NewReader(file)
	//readByLine(file)
	//按行读, 全部读出
	//readByLine(file)
	//
	////关闭文件句柄
	//fmt.Println()
	//defer file.Close()
	//fmt.Println("关闭文件成功....")
	//
	////3. 读取整个文件
	//readByFile()
}
