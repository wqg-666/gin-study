package main

import (
	"bytes"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//不同请求体对应的原始内容
	r.POST("", func(c *gin.Context) {
		//11111body格式
		byteData, _ := io.ReadAll(c.Request.Body)
		fmt.Println(string(byteData))
		//读了之后，body就没了，阅后即焚
		//解决body阅后即焚
		c.Request.Body = io.NopCloser(bytes.NewBuffer(byteData))
		//name := c.PostForm("name")
		//fmt.Println(name)
		//form-data
		//----------------------------771026781052387263253410//分隔符
		//Content-Disposition: form-data; name="name"
		//
		//琪琪
		//----------------------------771026781052387263253410//分隔符
		//Content-Disposition: form-data; name="age"
		//
		//123
		//----------------------------771026781052387263253410--//分隔符
		fmt.Println(c.Request.Header)
		//可以得到分隔符是
		//Content-Type:[multipart/form-data; boundary=--------------------------651187128909441928175877]
	})
	//2222x-www-form-urlencoded
	//url 为 name=%E7%90%AA%E7%90%AA&age=123
	//3333json格式
	//{
	//	"name":琪琪,
	//	"age":20
	//}
	r.Run(":8080")

}
