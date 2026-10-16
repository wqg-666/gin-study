package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	//111.用于浏览器直接请求这个接口唤起下载
	//r.GET("", func(c *gin.Context) {
	//	//1.要设置Content-Type，唤起浏览器下载，告诉客户端（浏览器）返回的数据类型：application/octet-stream`=二进制流
	//	//2.只能是GET请求
	//	//3.`Content-Disposition` 这个头专门控制文件展示方式：
	//	//- `attachment`：以附件形式，触发下载（关键！没有这个，浏览器可能直接打开内容而不是下载）
	//	//- `filename=3.响应文件.go`：浏览器下载弹窗里显示的文件名
	//	//4.读取服务器本地的文件 `3.响应文件.go`，把文件的二进制内容放到响应体里，发给浏览器。
	//	c.Header("Content-Type", "application/octet-stream")
	//	c.Header("Content-Disposition", "attachment; filename=3.响应文件.go")
	//	c.File("3.响应文件.go")
	//})
	//222.前端请求后端接口，然后唤起浏览器下载

	// 案例2的接口 /download2
	r.Static("st", "static")
	r.GET("/download2", func(c *gin.Context) {
		// 允许前端 JavaScript 读取自定义响应头
		c.Header("Access-Control-Expose-Headers", "fileName,msg")
		// 设置文件名
		c.Header("fileName", "xiaomao.png")
		// 设置文件类型
		c.Header("Content-Type", "image/png")
		// 返回图片文件
		c.File("uploads/xiaomao.png")
	})
	//333.前端唤起浏览器下载的本质,在这个static文件下
	//<a href="文件地址" download="文件名">文件下载</a>
	//最好的做法，点击static里的浏览器下载.html里的go图标
	//调下载接口的请求，后端不返回实际文件内容，而是生成一个临时下载地址
	//前端构造a标签，再请求这个接口唤起浏览器下载
	r.Run(":8080")
}
