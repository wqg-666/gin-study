package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	//?key=xxx&name=xxx&name=yyy,这种就被成为查询参数
	//查询参数不是GET请求专属的
	r := gin.Default()
	r.GET("", func(c *gin.Context) {
		//浏览器输入http://127.0.0.1:8080/?name=qiqi&age=123&key=123&key=124
		//控制台输出qiqi 123 [123 124]
		//http://127.0.0.1:8080/?name=qiqi&age=123&key=123
		//qiqi 123 [123]
		//只拿第一个key 的值，返回 string
		//name := c.Query("name")
		//	GET /?name=Manu&lastname=
		name, e := c.GetQuery("name")
		fmt.Printf("name:%s,e:%v\n", name, e) //name:Manu,e:true
		id, e := c.GetQuery("id")
		fmt.Printf("id:%s,e:%v\n", id, e) //id:,e:false
		lastname, e := c.GetQuery("lastname")
		fmt.Printf("lastname:%s,e:%v\n", lastname, e) //lastname:,e:true
		//age参数不存在时，默认值为25
		age := c.DefaultQuery("age", "25")
		//从URL查询参数，读取同名的多个参数，返回字符串切片 `[]string`。
		keyList := c.QueryArray("key")
		fmt.Println(name, age, keyList) //Manu 25 []
	})

	r.Run(":8080")
}
