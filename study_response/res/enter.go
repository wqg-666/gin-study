package res

import "github.com/gin-gonic/gin"

// 自定义的响应封装
// 自定义的响应结构体，规定所有接口返回json固定格式
type Response struct {
	Code int    `json:"code"` //业务码：0成功，非0代表各种错误
	Msg  string `json:"msg"`  //提示文字，给前端展示
	Data any    `json:"data"` //真实业务数据，成功时放内容，失败一般为nil
}

// 错误码映射表：key是业务code，value是对应的提示文字
var codeMap = map[int]string{
	1001: "权限错误",
	1002: "角色错误",
}

// response 底层核心私有函数（首字母小写，只能res包内部使用，外部main包不能直接调用）
func response(c *gin.Context, code int, data any, msg string) {
	//返回200状态码 + Response结构体，自动序列化成JSON
	c.JSON(200, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

// Ok 基础成功函数。自定义消息+返回数据
func Ok(c *gin.Context, data any, msg string) {
	//函数复用，函数嵌套调用
	response(c, 0, data, msg) //code固定传0，代表业务成功
}

// OkWithData,成功，只传data，消息默认"成功"
func OkWithData(c *gin.Context, data any) {
	Ok(c, data, "成功")
}

// OkWithMsg：成功，自定义提示消息，data传空map{}
func OkWithMsg(c *gin.Context, msg string) {
	Ok(c, map[string]any{}, msg)
}

// Fail 基础失败函数，可以自定义错误code、data、msg
func Fail(c *gin.Context, code int, data any, msg string) {
	response(c, code, data, msg)
}

// FailWithMsg：快速失败，自定义提示文字，默认业务码1001，data=nil
func FailWithMsg(c *gin.Context, msg string) {
	response(c, 1001, nil, msg)
}

// FailWithCode：只传业务码，自动从codeMap里读取对应的msg；找不到则提示"服务错误"
func FailWithCode(c *gin.Context, code int) {
	msg, ok := codeMap[code]
	if !ok {
		msg = "服务错误"
	}
	response(c, code, nil, msg)
}
