package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

//客户关系管理系统欢迎界面函数
func wel() {
	wel := `
---------------------客户信息管理系统---------------------
		1. 添加客户
		2. 查看客户
		3. 更新客户
		4. 删除客户
		5. 保存客户信息
		6. 查看已保存的客户信息
		7. 退出客户信息管理系统
------------------------------------------------------------
`
	fmt.Println(wel)
}

//客户信息管理系统Switch函数
func (c CmsData) choice() {
	for true {
		var choice string
		wel()
		fmt.Println("请输入您的选择[1-7]: ")
		fmt.Scan(&choice)
		switch choice {
		case "1":
			{
				//添加客户信息功能
				c.addCms()
			}
		case "2":
			{
				//查询客户信息功能
				c.chaCms()
			}
		case "3":
			{
				for true {
					//修改客户信息功能
					//fmt.Println("请输入要修改的客户信息编号:")
					//var id int
					//fmt.Scan(&id)
					cmsid := cmsID()
					c.modfCms(cmsid - 1)
					b := isBack()
					if b {
						break
					}
				}
			}
		case "4":
			{
				//删除客户信息功能
				c.deleteCms()
			}
		case "5":
			{
				//保存客户信息
				c.cunCms()
			}
		case "6":
			{
				//查看已保存的客户信息
				c.chaJsonCmsData()
			}
		case "7":
			{
				//退出客户信息管理系统
				fmt.Println("您已退出客户信息管理系统,欢迎再次使用...")
				break
			}
		default:
			{
				//非法输入
				fmt.Println("您输入的选择数值不合法,请重新输入,,,")
				continue
			}
		}
	}
}

//返回上一层函数
func isBack() bool {
	var back string
	fmt.Println("请问是否返回上一层[Y/N]: ")
	fmt.Scan(&back)
	if strings.ToUpper(back) == "Y" {
		return true
	} else {
		return false
	}
}

//定义客户信息数据结构
type CmsInfo struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int8   `json:"age"`
	Email  string `json:"email"`
}

//客户信息构造函数
func NewCmsInfo(name string, gender string, age int8, email string) *CmsInfo {
	return &CmsInfo{
		Name:   name,
		Gender: gender,
		Age:    age,
		Email:  email,
	}
}

type CmsData struct {
	cmsdata []CmsInfo
}

func NewCmsData(data []CmsInfo) *CmsData {
	return &CmsData{
		cmsdata: data,
	}
}

func (n *CmsData) CmsID() (int, *CmsInfo) {
	var sid int
	idCMS := &(n.cmsdata[sid-1])
	return sid, idCMS
}

//输入客户信息函数
func inputCms() (name, gender string, age int8, email string) {
	fmt.Println("请输入客户姓名: ")
	fmt.Scan(&name)
	fmt.Println("请输入客户性别: ")
	fmt.Scan(&gender)
	fmt.Println("请输入客户年龄:")
	fmt.Scan(&age)
	fmt.Println("请输入客户邮箱: ")
	fmt.Scan(&email)
	return name, gender, age, email
}

//添加客户信息功能函数
func (c *CmsData) addCms() {
	for true {
		fmt.Println("-------------------添加客户信息开始-------------------")
		name, gender, age, email := inputCms()
		newCmsdata1 := NewCmsInfo(name, gender, age, email)
		c.cmsdata = append(c.cmsdata, *newCmsdata1)
		//fmt.Println(c)
		fmt.Println("-------------------添加客户信息结束-------------------")
		b := isBack()
		if b {
			break
		}
	}
}

//查询客户信息函数
func (c *CmsData) chaCms() {
	for true {
		fmt.Println("-------------------查询客户信息开始-------------------")
		for i, v := range c.cmsdata {
			fmt.Printf("编号:%-8d 姓名:%-8s 性别:%-8s 年龄:%-8d 邮箱:%-8s\n", i+1, v.Name, v.Gender, v.Age, v.Email)
		}
		fmt.Println("-------------------查询客户信息结束-------------------")
		b := isBack()
		if b {
			break
		}
	}

}

//客户信息编号函数
func cmsID() int {
	var Sid int
	fmt.Println("请输入要操作的客户信息编号:")
	fmt.Scan(&Sid)
	return Sid
}

//更新客户信息函数
func (c *CmsData) modfCms(id int) {

	fmt.Println("-------------------修改客户信息开始-------------------")
	fmt.Println("请输入要修改的客户姓名: ")
	fmt.Scan(&c.cmsdata[id].Name)
	fmt.Println("请输入要修改的客户性别: ")
	fmt.Scan(&c.cmsdata[id].Gender)
	fmt.Println("请输入要修改的客户年龄: ")
	fmt.Scan(&c.cmsdata[id].Age)
	fmt.Println("请输入要修改的客户邮箱: ")
	fmt.Scan(&c.cmsdata[id].Email)
	fmt.Println("-------------------修改客户信息结束-------------------")

}

/*//删除客户信息函数
func (c *CmsData) deleteCms() {
	for true {
		fmt.Println("-------------------删除客户信息开始-------------------")
		var deleteId int
		fmt.Println("请输入要删除的客户信息编号:")
		fmt.Scan(&deleteId)
		if deleteId == -1 {
			fmt.Println("您输入的客户信息编号不存在,请重新输入,,,")
			continue
		}
		deleteId = deleteId - 1
		c.cmsdata = append(c.cmsdata[:deleteId], c.cmsdata[deleteId+1:]...)
		fmt.Println("-------------------删除客户信息结束-------------------")
		b := isBack()
		if b {
			break
		}
	}
}*/

//删除客户信息函数
func (c *CmsData) deleteCms() {
	for true {
		fmt.Println("-------------------删除客户信息开始-------------------")
		cmsid := cmsID()
		if cmsid == -1 {
			fmt.Println("您输入的客户信息编号不存在,请重新输入,,,")
			continue
		}
		cmsid = cmsid - 1
		c.cmsdata = append(c.cmsdata[:cmsid], c.cmsdata[cmsid+1:]...)
		fmt.Println("-------------------删除客户信息结束-------------------")
		b := isBack()
		if b {
			break
		}
	}
}

//保存客户信息函数
func (c *CmsData) cunCms() {
	for true {
		fmt.Println("-------------------保存客户信息开始-------------------")
		jsonCmsData, err := json.Marshal(c.cmsdata)
		//fmt.Println(string(jsonCmsData))
		if err != nil {
			fmt.Println("客户信息保存失败xxx")
		} else {
			fmt.Println("客户信息保存成功***")
		}
		ioutil.WriteFile("jsonCmsData.json", jsonCmsData, 0666)
		fmt.Println("-------------------保存客户信息结束-------------------")
		b := isBack()
		if b {
			break
		}
	}
}

//查询已保存的客户信息功能函数
func (c *CmsData) chaJsonCmsData() {
	for true {
		fmt.Println("-----------------查看已保存客户信息开始-----------------")
		cmsData1, err := ioutil.ReadFile("/Users/denis/code/golang_study/src/jsonCmsData.json")
		if err != nil {
			fmt.Println("打开已保存客户信息失败xxx")
		} else {
			fmt.Println("打开已保存客户信息成功****")
		}
		fmt.Println(string(cmsData1))
		fmt.Println("-----------------查看已保存客户信息结束-----------------")
		b := isBack()
		if b {
			break
		}
	}

}

func main() {
	//wel()

	newCmsDataInfo := NewCmsData([]CmsInfo{})
	newCmsDataInfo.choice()

}
