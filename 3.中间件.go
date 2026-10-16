package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	fmt.Println("Home")
	fmt.Println(c.Get("GM1")) //输出true
	fmt.Println(c.Get("GM2")) //输出true
	//`c.Get("user")`：从gin上下文里，读取 key 为`user`的值
	_user, ok := c.Get("user")
	//给客户端返回文本响应`Home`，状态码 200。
	c.String(200, "Home")
	if ok {
		//因为`c.Get`返回的是万能空接口`interface{}`，不知道里面存的真实类型是什么，需要断言
		//断言`_user`里面存的数据就是`UserInfo`结构体类型
		user, ok := _user.(UserInfo)
		//`user`就是 UserInfo 结构体对象，打印结构体里的 Name 字段。
		if ok {
			fmt.Println(user.Name)
		}
	}
}

// 局部中间件，作用于单个路由
func M1(c *gin.Context) {
	fmt.Println("M1 请求部分")
	c.Abort() //拦截
	c.String(200, "M1返回的响应")
	//fmt.Println("M1 响应部分")
}
func M2(c *gin.Context) {
	fmt.Println("M2 请求部分")
	c.Next() //通行
	fmt.Println("M2 响应部分")
}

type UserInfo struct {
	Name string
}

// 全局中间件，作用于一组路由
func GM1(c *gin.Context) {
	fmt.Println("GM1 请求部分")
	var user = UserInfo{Name: "qq"}
	//c.Set,中间件传参数
	c.Set("GM1", "GM1")
	c.Set("user", user)
	c.Next() //通行
	fmt.Println("GM1 响应部分")
}
func GM2(c *gin.Context) {
	fmt.Println("GM2 请求部分")
	c.Set("GM2", "GM2")
	c.Next() //通行
	fmt.Println("GM2 响应部分")
}

func main() {
	r := gin.Default()
	//Default,包含有Recovery()，函数中有panic()程序不会停掉
	g := r.Group("api")
	g.Use(GM1) //走GM1的请求，走GM1的响应
	g.Use(GM2) //走GM2的请求，走GM2的响应
	//或者
	//g.Use(GM1, GM2)
	g.GET("users", Home) //走到Home
	//全局中间件
	//GM1 请求部分
	//GM2 请求部分
	//Home
	//GM2 响应部分
	//GM1 响应部分

	//局部中间件
	//r.GET("/", M1, M2, Home)
	//中间件M1,M2，视图为Home
	//使用Abort()后会原路返回
	//M1,M2 c.next()
	//M1 请求部分
	//M2 请求部分
	//Home
	//M2 响应部分
	//M1 响应部分

	//M1,c.Abort() M2,c.next()
	//M1 请求部分
	//M1 响应部分

	//M1,c.next() M2,c.Abort()
	//M1 请求部分
	//M2 请求部分
	//M2 响应部分
	//M1 响应部分

	//全局中间件传参数的结果
	//GM1 请求部分
	//GM2 请求部分
	//Home
	//GM1 true
	//GM2 true
	//qq
	//GM2 响应部分
	//GM1 响应部分
	r.Run(":8080")
}
