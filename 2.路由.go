package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func UserView(c *gin.Context) {
	//`c *gin.Context`：Gin 上下文，最重要的对象，封装了本次 HTTP 请求全部信息：请求参数、请求头、url、响应写入等。
	//一次请求对应一个`Context`。
	//`c.Request.URL`：拿到本次请求的完整 URL 对象。
	//`c.Request.Method`：获取 HTTP 请求方法（`GET/POST/PUT/DELETE`）
	path := c.Request.URL
	fmt.Println(c.Request.Method, path)
}
func main() {
	r := gin.Default()
	//`r.Group("api")`：路由分组，创建前缀为`/api`的路由组。
	//组下所有路由都会自动拼接前缀。这个组里的路由完整路径都会以 `/api` 开头
	apiGroup := r.Group("api")
	//`serGroup(apiGroup)`：把这个`apiGroup`传给`UserGroup`函数，在里面注册 users 相关路由。
	UserGroup(apiGroup)
	//又新建一个路由组，前缀同样是`/api`，名字叫`noMiddleWareGroup`（无中间件组）
	//交给`LoginGroup`注册 login 相关路由
	noMiddleWareGroup := r.Group("api")
	LoginGroup(noMiddleWareGroup)
	r.Run(":8080")
}

func UserGroup(r *gin.RouterGroup) {
	r.GET("users", UserView)
	r.POST("users", UserView)
	r.DELETE("users", UserView)
	r.PUT("users", UserView)
}
func LoginGroup(r *gin.RouterGroup) {
	r.GET("login", UserView)
	r.POST("login", UserView)
	r.DELETE("login", UserView)
	r.PUT("login", UserView)
}

//RESETFul Api规范,尽量使用名词的复数来定义路由
//只是规范，不是要求
//GET /api/users 用户创建
//POST /api/users 创建用户
//PUT /api/users/:id 更新用户信息
//DELETE /api/users/:id 删除单个用户
//DELETE /api/users 批量删除用户

//大家在公司里面的项目，基本上都是POST请求
//很早之前的，那个时候还没有RESETFul 规范这个说法
//很多公司的防火墙会拦截除GET和POST之外的请求
