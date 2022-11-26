package core

import "github.com/gin-gonic/gin"

func Book(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "查看所有书籍",
	})
}

func AddBook(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "添加书籍",
	})
}

func EditBook(c *gin.Context) {
	//删除书籍功能
	c.JSON(200, gin.H{
		"message": "编辑书籍",
	})
}

func DeleteBook(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "删除书籍",
	})
}
