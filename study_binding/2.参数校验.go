package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("json", func(c *gin.Context) {
		type User struct {
			//如果有多个规则使用用逗号分割
			//Name  string `json:"name" binding:"required,min=3,max=5"`
			//Age   int    `json:"age"`
			//Pwd   string `json:"pwd" binding:"required"`
			//RePwd string `json:"rePwd" binding:"eqfield=Pwd"`
			//IP string `json:"ip" binding:"required,ip"`
			//对数组IPList里的每个元素校验是不是为空，是不是ip地址，不能校验空数组
			IPList []string `json:"iPList" binding:"dive,required,ip"`
		}
		var user User
		err := c.ShouldBindJSON(&user)
		if err != nil {
			c.String(200, err.Error())
			return
		}
		c.JSON(200, user)
		//传入空name,Key: 'User.Name' Error:Field validation for 'Name' failed on the 'required' tag
		//不传name,Key: 'User.Name' Error:Field validation for 'Name' failed on the 'required' tag
	})
	r.Run()
}

//|/required：必填字段，如：binding:"required"
//min 最小长度，如：binding:"min=5"
//max最大长度，如：binding:"max=10"
//len 长度，如：binding:"len=6"
//|/针对数字的大小
//eq 等于，如：binding:"eq=3"
//ne不等于，如：binding:"ne=12"
//gt 大于，如：binding:"gt=1@"
//gte 大于等于，如：binding:"gte=1@"
//lt 小于，如：binding:"Lt=10"
//lte 小于等于，如：binding:"lte=10"
//|/针对同级字段的
//eqfield 等于其他字段的值，如：PassWord string `binding:"eqfield=Passwordi
//nefield 不等于其他字段的值
//忽略字段，如：binding:"-”或者不写
//|/枚举只能是red 或green
//oneof=red green
//|/字符串
//contains=fengfeng //包含fengfeng的字符串
//excludes //不包含
//startswith |/ 字符串前缀
//endswith |/ 字符串后缀
//|/数组
//dive // dive后面的验证就是针对数组中的每一个元素
//|/网络验证
//ip
//ipv4
//ipv6
//uri
//url
//// uri 在于I(Identifier)是统一资源标示符，可以唯一标识一个资源。
//// urz 在于Locater，是统一资源定位符，提供找到该资源的确切路径
////日期验证1月2号下午3点4分5秒在2006年
//datetime=2006-01-02

//## 一、`binding:"url"` 校验（完整网址）
//要求：必须包含协议 `http://` 或者 `https://`
//type Test struct {
//    Link string `json:"link" binding:"required,url"`
//✅ 通过
//{"link":"https://www.baidu.com"}
//{"link":"http://127.0.0.1:8080/test"}
//❌ 失败
//{"link":"www.baidu.com"} //缺少http协议，url校验不通过
//{"link":"baidu"}

//## 二、`binding:"uri"` 校验（资源标识符，可以不带协议）
//范围更广！URL 属于 URI 子集。可以是相对路径、不带协议的地址
//type Test struct {
//    Resource string `json:"resource" binding:"required,uri"`
//}
//✅ 通过
//{"resource":"/user/list"}
//{"resource":"https://www.baidu.com"}
//{"resource":"user/1001"}
// URI（统一资源标识符 Identifier）：用来唯一标识资源，只负责 “给资源一个编号”，不一定能定位。
// URL（统一资源定位符 Locator）：属于 URI 的子集，既标识资源，又给出访问地址（在哪里）
