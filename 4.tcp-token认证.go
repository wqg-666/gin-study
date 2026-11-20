package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type User struct {
	Username string
	Password string
}

var (
	// Gin 会并发处理请求，因此需要保护共享数据。
	stateMu sync.RWMutex

	users = []User{
		{
			Username: "zhangsan",
			Password: "123456",
		},
	}

	// token -> username
	tokenUsers = make(map[string]string)
)

// 登录
func login(ctx *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "请求参数错误",
		})
		return
	}

	// 查找匹配的用户
	username := ""

	stateMu.RLock()
	for _, user := range users {
		if user.Username == req.Username &&
			user.Password == req.Password {
			username = user.Username
			break
		}
	}
	stateMu.RUnlock()

	if username == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "账号或密码错误",
		})
		return
	}

	// 生成随机 Token
	token := uuid.NewString()

	// 保存 Token 与用户的对应关系
	stateMu.Lock()
	tokenUsers[token] = username
	stateMu.Unlock()

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"token":   token,
	})
}

// 从 Authorization 请求头中解析 Bearer Token
func parseBearerToken(header string) (string, bool) {
	scheme, token, ok := strings.Cut(header, " ")
	if !ok || !strings.EqualFold(scheme, "Bearer") {
		return "", false
	}

	if token == "" {
		return "", false
	}

	return token, true
}

// 身份认证中间件
func authRequired(ctx *gin.Context) {
	token, ok := parseBearerToken(
		ctx.GetHeader("Authorization"),
	)
	if !ok {
		ctx.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{"message": "认证失败"},
		)
		return
	}

	// 根据 Token 查找用户
	stateMu.RLock()
	username, exists := tokenUsers[token]
	stateMu.RUnlock()

	if !exists {
		ctx.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{"message": "认证失败"},
		)
		return
	}

	// 将用户信息保存到本次请求的 Context 中
	ctx.Set("username", username)

	// 继续执行后面的处理函数
	ctx.Next()
}

// 修改当前用户的密码
func updatePassword(ctx *gin.Context) {
	//要从前端写入的东西
	var req struct {
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "请求参数错误",
		})
		return
	}

	// 获取认证中间件（authRequired）保存的用户名
	username := ctx.GetString("username")

	updated := false

	stateMu.Lock()
	for index := range users {
		if users[index].Username == username {
			users[index].Password = req.Password
			updated = true
			break
		}
	}
	stateMu.Unlock()

	if !updated {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "用户不存在",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "密码修改成功",
	})
}

func main() {
	router := gin.New()

	// 使用 Use 添加中间件，它们会应用到之后注册的所有路由
	// gin.Logger() 是 Gin 内置的请求日志中间件，类似前面写的 requestLogger
	router.Use(gin.Logger())

	// gin.Recovery() 是 Gin 内置的 panic 恢复中间件。
	// 避免整个 HTTP 服务因为当前请求的 panic 而退出
	router.Use(gin.Recovery())

	// 登录接口不需要认证
	router.POST("/login", login)

	// /api 下的接口都需要认证
	api := router.Group("/api")
	//代码跑到这里，`authRequired` 函数本身不会执行，只是告诉 Gin：
	//以后只要有人请求 `/api/xxx` 下面的接口，先跑 `authRequired`，跑完没问题，再去执行后面的处理函数（`updatePassword`）
	api.Use(authRequired) //中间件，所以属于api的组都进行了authRequired
	{
		api.PATCH("/users/me/password", updatePassword)
	}

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
