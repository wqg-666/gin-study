package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("users", func(c *gin.Context) {
		name := c.PostForm("name")
		age, ok := c.GetPostForm("age")
		fmt.Println(name)
		fmt.Println(age, ok)
		//ok为true代表传数据了，可能为空，为false表示没有传age这个参数
		//在apifox中Body中的form-data中加入name,age参数，然后发送
		//输出琪琪
		//123 true
	})
	r.Run(":8080")
}
