package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

// 全局变量：翻译器实例，用来把校验错误翻译成中文。
var trans ut.Translator

// `init`函数在`main`之前自动运行，只执行一次。
func init() {
	//1111创建中文翻译器
	uni := ut.New(zh.New())
	trans, _ = uni.GetTranslator("zh")

	//1111注册中文翻译器到Gin的校验引擎
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		_ = zh_translations.RegisterDefaultTranslations(v, trans)
	}
	//2222字段名也显示中文
	//return 出去的字符串，会给到 validator，validator 拿这个字符串，作为错误信息里的字段名称
	v.RegisterTagNameFunc(func(field reflect.StructField) string {
		//读取结构体字段 `tag` 中 `label` 的值。
		label := field.Tag.Get("label")
		//如果 label 为空，返回原生结构体字段名 `field.Name`
		if label == "" {
			return field.Name
		}
		//2222的返回，return label
		//3333错误信息和错误字段一起返回
		//{
		// "name": "用户姓名为必填字段"
		//}
		name := field.Tag.Get("json")
		//3333的返回
		return fmt.Sprintf("%s---%s", name, label)
		//`fmt.Printf()`：格式化并且直接打印到控制台，返回打印字节数
		//`fmt.Sprintf()`：格式化，生成字符串返回，不输出
		//Name---用户姓名

	})
}

// `RegisterTagNameFunc` 的返回值，会被 validator 保存，在 `e.Field()` 里面拿到
func ValidateErr(err error) any {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return err.Error()
	}
	//判断错误是不是validator校验错误
	//1. 如果是，循环调用`e.Translate(trans)`翻译成中文
	//2. 多条错误用`;`拼接成一句话返回 示例：`Name为必填字段;Email为必填字段`
	//1111.2222
	//var list []string
	//for _, e := range errs {
	//	list = append(list, e.Translate(trans))
	//}
	//return strings.Join(list, ";")
	//3333
	//map key：字符串（前端 JSON 字段名，比如`name`）
	//value：任意类型，可以放字符串、数字等，这里放错误文本
	var m = map[string]any{}
	for _, e := range errs {
		msg := e.Translate(trans)
		//按照---切割,"name---用户姓名为必填字段"
		_list := strings.Split(msg, "---")
		m[_list[0]] = _list[1]
		//m["name"] = "用户姓名为必填字段"
	}
	return m
}

type User struct {
	Name  string `json:"name" binding:"required" label:"用户姓名"`
	Email string `json:"email" binding:"required,email" label:"邮箱"`
}

func main() {
	r := gin.Default()
	// 注册路由
	r.POST("/user", func(c *gin.Context) {
		var user User
		//从请求体JSON绑定到user结构体，并自动执行binding校验
		if err := c.ShouldBindJSON(&user); err != nil {
			// 参数验证失败，返回中文错误字符串1111,2222
			//c.String(200, ValidateErr(err))
			//3333
			c.JSON(200, map[string]any{
				"code": 7,
				"msg":  "验证错误",
				"data": ValidateErr(err), //为返回的var m = map[string]any{}
			})
			//1111 Email必须是一个有效的邮箱
			//2222 用户姓名为必填字段;邮箱必须是一个有效的邮箱
			//3333
			//{
			//    "code": 7,
			//    "data": {
			//        "email": "邮箱必须是一个有效的邮箱"
			//    },
			//    "msg": "验证错误"
			//}
			return
		}

		// 参数验证成功，返回JSON响应
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Hello, %s! Your email is %s.", user.Name, user.Email),
		})
		//1111 2222
		//{
		//    "message": "Hello, ui! Your email is 3034979534@qq.com."
		//}

	})

	// 启动HTTP服务器
	r.Run()
}
