package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	//`r.Static("st", "static")`：映射一整个文件夹，可以访问文件夹里所有文件
	//`r.StaticFile("abc", "static/abc.txt")`：只映射单独一个文件，只绑定这一条 URL
	//http://127.0.0.1:8080/st/abc.txt,访问的地址
	//http://127.0.0.1:8080/st/浏览器下载.html
	//第一个参数 `"st"`：路由前缀（URL 路径）
	//第二个参数 `"static"`：服务器本地文件夹名
	r.Static("st", "static")

	//第一个参数：URL 路由地址：`abc`，也就是访问 `/abc`
	//第二个参数：服务器本地文件路径：`static/abc.txt`
	r.StaticFile("abc", "static/abc.txt")
	//静态文件的路径不能再被路由使用，这两个的地址重复了，都是abc
	r.GET("abc", func(c *gin.Context) {
		c.JSON(200, gin.H{})
	})
	r.Run(":8080")
}
