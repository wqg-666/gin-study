package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//`err != nil` 只有两大类原因：
	//1. 解析失败：拿到了参数的值，但是转不成结构体字段的类型（数据本身格式不对）
	//2. 校验不通过：类型转换成功了，但是不满足你写的 `binding` 规则（数值合法，但业务规则不满足）
	//查询参数,Params是放数据的位置
	r.GET("", func(c *gin.Context) {
		type User struct {
			Name string `form:"name"`
			Age  int    `form:"age"`
		}
		var user User
		err := c.ShouldBindQuery(&user)
		fmt.Println(user, err)
		//{qq 123} <nil>
		//{qq 0} strconv.ParseInt: parsing "\"111\"": invalid syntax

	})
	//路径参数,动态参数，Params是放数据的位置
	r.GET("users/:id/:name", func(c *gin.Context) {
		type User struct {
			Name string `uri:"name"`
			ID   int    `uri:"id"`
		}
		var user User
		err := c.ShouldBindUri(&user)
		fmt.Println(user, err)
		//{qiqi 1} <nil>
	})
	//表单参数，Body请求体form-data是放参数的位置
	r.POST("form", func(c *gin.Context) {
		type User struct {
			Name string `form:"name"`
			Age  int    `form:"age"`
		}
		var user User
		err := c.ShouldBind(&user)
		fmt.Println(user, err)
		//{琪琪 123} <nil>
	})
	//JSON参数，Body请求体json是放参数的位置
	r.POST("json", func(c *gin.Context) {
		type User struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}
		var user User
		err := c.ShouldBindJSON(&user)
		fmt.Println(user, err)
		//{琪琪 20} <nil>
	})
	//Header参数,Headers是放参数的位置
	r.POST("header", func(c *gin.Context) {
		type User struct {
			Name        string `header:"Name"`
			Age         int    `header:"Age"`
			UserAgent   string `header:"User-Agent"`
			ContentType string `header:"Content-Type"`
		}
		var user User
		err := c.ShouldBindHeader(&user)
		fmt.Println(user, err)
		//{qiqi 20 Apifox/1.0.0 (https://apifox.com) application/json} <nil>
	})
	r.Run(":8080")
}

//GET 和 POST 的核心区别
//1.GET
//1. 习惯：数据放在 URL 后面 Query 参数（`?name=qq`）
//2. URL 有长度限制，不能传大量数据
//3. 一般用来查询数据（读操作，不会修改服务器数据）
//4. 浏览器会缓存 GET 请求；URL 参数会保存在浏览器历史记录里，不能传密码、敏感信息
//5. 没有请求体（行业规范，虽然协议底层技术上可以加 body，但浏览器不支持）
//## 2.POST
//1. 习惯：数据放在请求体 Body 里面（json /form 表单，适合大量数据）
//2. 没有 URL 长度限制，可以传大内容（比如上传图片、复杂 JSON）
//3. 一般用来新增 / 提交数据（写操作，修改服务器资源）
//4. 不会默认缓存，参数不会出现在浏览器地址栏，适合提交表单
