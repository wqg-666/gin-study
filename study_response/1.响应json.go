package main

import (
	"gin_study/study_response/res"

	"github.com/gin-gonic/gin"
)

func main() {
	//1.初始化
	r := gin.Default()
	//2.挂载路由
	//`/index`：访问路径，完整地址 `http://127.0.0.1:8080/index`
	//状态码200。`200`：请求正常，直接读取返回的 JSON 数据（code、msg、data）
	//- `404`： 找不到这个接口路径 `500`： 服务器内部出错 panic
	//`code` 是我们自己定义的业务标记，用来描述接口里面业务逻辑成功还是失败，
	//- `code = 0` → 业务成功 `code = 1001` → 用户名或密码错误
	//`code = 1002` → 账号不存在 `code = 1003` → token 失效，请重新登录
	r.GET("/login", func(c *gin.Context) {
		//c.JSON(200, gin.H{"code": 0, "msg": "成功", "data": gin.H{}})
		res.OkWithMsg(c, "登录成功")
	})
	r.GET("/users", func(c *gin.Context) {
		//c.JSON(200, gin.H{"code": 0, "msg": "成功", "data": gin.H{}})
		res.OkWithData(c, map[string]any{
			"name": "琪琪",
		})
	})
	r.POST("/users", func(c *gin.Context) {
		//c.JSON(200, gin.H{"code": 0, "msg": "成功", "data": gin.H{}})
		res.FailWithMsg(c, "参数错误")
	})
	//3.绑定端口
	r.Run(":8080")
}
