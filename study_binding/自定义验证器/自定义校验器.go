package main

import (
	"fmt"
	"net"
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

var trans ut.Translator

func init() {
	uni := ut.New(zh.New())
	trans, _ = uni.GetTranslator("zh")
	//`binding.Validator.Engine()`：拿到 Gin 底层使用的 validator 校验引擎
	// 类型断言转成`*validator.Validate`指针
	v, ok := binding.Validator.Engine().(*validator.Validate)
	//如果拿到 validator 实例，注册默认中文翻译，
	//validator自带规则 (required/min 等) 错误信息会变成中文。`_`忽略返回的 error。
	if ok {
		_ = zh_translations.RegisterDefaultTranslations(v, trans)
	}
	//自定义读取结构体标签的函数
	v.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		if label == "" {
			return field.Name
		}
		name := field.Tag.Get("json")
		return fmt.Sprintf("%s---%s", name, label)
		//返回ip---ip地址
	})
	//注册自定义校验规则，tag 名字叫`fip`。当结构体字段写`binding:"fip"`就会执行这个函数。
	//`fl validator.FieldLevel`：当前字段的上下文对象，可以获取字段值、参数、结构体信息
	v.RegisterValidation("fip", func(fl validator.FieldLevel) bool {
		//返回值`bool`：`true=校验通过；false=校验失败`
		fmt.Println("fl.Field():", fl.Field())                    //反射 Value，字段的值
		fmt.Println("fl.FieldName()", fl.FieldName())             //自定义 tag 返回的名称（我们前面返回的`ip---ip地址`）
		fmt.Println("fl.StructFieldName()", fl.StructFieldName()) //结构体原始字段名`Ip`
		fmt.Println("fl.Parent()", fl.Parent())                   //父结构体（嵌套结构体才有用，这里就是 User）
		fmt.Println("fl.Top():", fl.Top())                        //最顶层结构体
		fmt.Println("fl.Param()", fl.Param())                     //tag 参数，比如`binding:"fip=123"`，Param 拿到`123`，
		//`.(string)` —— 类型断言，尝试把 `interface{}` 里面的值，强制断言成 `string` 类型。
		ip, ok := fl.Field().Interface().(string)
		//如果是字符串，并且 ip 不为空
		if ok && ip != "" {
			ipObj := net.ParseIP(ip)
			return ipObj != nil //合法返回 true，不合法返回 false
		}
		return true //如果 ip 是空字符串`""`，直接`return true`
		//这个回调函数的返回值是给 validator 库内部用的，不是你代码里直接接收。
		//`return true` → validator：此字段校验没问题，不产生错误
		//`return false` → validator：校验失败！生成一条 `validator.ValidationError` 错误记录
	})
}

// 自定义错误处理函数，把 validator 校验错误，加工成前端友好的 map 格式。
func ValidateErr(err error) any {
	errs, ok := err.(validator.ValidationErrors)
	//判断 err 是不是 validator 的校验错误。
	//如果是 validator 校验错误，进入下面处理。
	//validator 校验错误（字段规则不满足，如 ip 格式错）
	//类型：`validator.ValidationErrors`，是多条错误的集合，可以遍历、拿到 Tag、字段名、翻译信息
	///如果不是（比如 json 格式解析错误），直接返回原生错误字符串。
	//非 validator 错误（JSON 解析错误）
	//比如前端传了非法 json `{ip:123`（少大括号），这不是参数校验规则问题，是 json 解析失败。
	if !ok {
		return err.Error()
	}
	var m = map[string]any{}
	for _, e := range errs {
		//`e.Translate(trans)` 使用翻译器翻译错误信息，这里 msg 是我们前面自定义返回的`ip---ip地址`
		msg := e.Translate(trans)
		//得到切片`["ip", "ip地址"]`
		_list := strings.Split(msg, "---")
		//判断错误 tag 是不是我们自定义的`fip`
		//误区：只要开了中文翻译，所有 tag 自动翻译
		//真相：翻译是一张映射表，只有提前注册过的 tag 才有翻译。
		//内置 tag 已经由`zh_translations`帮你批量注册；
		//自己新增的自定义 tag，需要手动注册翻译文案，否则只能返回英文兜底文本。
		if e.Tag() == "fip" {
			//`e.Field()` → `ip---ip地址`
			m[strings.Split(e.Field(), "---")[0]] = "该ip地址不符合要求"
			//m[ip]="该ip地址不符合要求"
			continue //跳过m[_list[0]] = _list[1]
		}
		m[_list[0]] = _list[1]
	}
	return m
}

func main() {
	r := gin.Default()
	r.POST("/user", func(c *gin.Context) {
		type User struct {
			//`binding:"fip"`：启用我们注册的自定义校验规则 fip
			Ip string `json:"ip" binding:"fip" label:"ip地址"`
		}
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(200, map[string]any{
				"code": 7,
				"msg":  "验证错误",
				"data": ValidateErr(err),
			})
			return
		}
		c.JSON(http.StatusOK, user)
	})
	r.Run(":8080")
}
