package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//动态参数
	//用户的个人信息页面，他的路径
	//users?id=123 ,查询参数模式
	//users/123 ，动态参数模式
	r.GET("users/:id", func(c *gin.Context) {
		//浏览器输入http://127.0.0.1:8080/users/123
		//控制台输出123
		userID := c.Param("id")
		fmt.Println(userID)
	})
	r.GET("users/:id/:name", func(c *gin.Context) {
		//浏览器输入http://127.0.0.1:8080/users/123/qiqi
		//控制台输出123 qiqi
		userID := c.Param("id")
		userName := c.Param("name")
		fmt.Println(userID, userName)
	})
	//关于接口测试工具，postman,apifox
	//注意：接口测试工具能走通的，前端请求不一定可以走通
	//1.get请求带请求体
	//2.ws加请求头
	r.Run(":8080")
}
