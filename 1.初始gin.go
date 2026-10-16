package main

import "github.com/gin-gonic/gin"

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Index(c *gin.Context) {
	//法一
	c.JSON(200, Response{
		Code: 0,
		Msg:  "成功",
		Data: map[string]any{},
	})
	//法二
	//c.JSON(200, gin.H{})
	//Gin 提供的方法，直接向客户端返回JSON格式响应,`200`：HTTP 状态码 200，表示请求成功
	//`gin.H{}`：`gin.H` 本质是 `map[string]any`,用来快速构造 json 对象。
	//这里是空 map，所以返回空 JSON 对象 `{}`
}
func main() {
	//关闭debug输出
	gin.SetMode("release")
	//1.初始化
	//创建一个 Gin 的路由引擎实例（Engine），赋值给变量 `r`。
	r := gin.Default()
	//2.挂载路由
	r.GET("/index", Index)
	//3.绑定端口，运行
	r.Run(":8080") //等于r.Run("0.0.0.0:8080"),内网（同一个路由器下面的所有设备组成的小网络）的伙伴访问
	//r.Run("127.0.0.1:8080),本机访问
}
