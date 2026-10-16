package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// /*templates目录下的所有文件
	r.LoadHTMLGlob("study_response/templates/*")
	// 加载单个文件
	r.LoadHTMLFiles("study_response/templates/index.html")
	// 注册路由：访问【根路径】
	r.GET("", func(c *gin.Context) {
		//返回html页面
		//- `200`：HTTP 成功状态码
		//- `"index.html"`：指定要渲染的模板名字（就是前面加载的 html 文件名）
		//- `nil`：不往 html 页面传递任何变量数据
		//后端返回模板，可以很方便的修改网站标题和logo，服务端渲染网页
		c.HTML(200, "index.html", map[string]any{
			"title": "wqgnihao",
		})
	})
	r.Run(":8080")
}
