package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.POST("users", func(c *gin.Context) {
		name := c.PostForm("name")
		age, ok := c.GetPostForm("age")
		//`ok == true`：请求里携带了这个字段（哪怕值是空字符串）
		//`ok == false`：**请求里根本没有这个字段**
		//可以区分「没传参数」和「传了空参数」
		fmt.Println(name)
		fmt.Println(age, ok)
	})
	r.Run(":8080")
}
