package main

import (
	"fmt"
	"io"
	"net/http"
)

// `writer http.ResponseWriter`：响应写入器，你往这里写内容，内容就会发给浏览器 / Apifox（返回给客户端）
// `request *http.Request`：客户端发来的全部请求信息（请求方法、地址、请求头、提交的 body 数据都在这里）
func Index(writer http.ResponseWriter, request *http.Request) {
	//请求部分
	//- `request.Method`：请求方法，`GET` / `POST` / `PUT` / `DELETE`
	//- `request.URL.String()`：请求的地址，比如 `http://127.0.0.1:8080/index`
	fmt.Println(request.Method, request.URL.String())
	//GET 请求**没有 Body**，所以不需要读取，直接跳过这段代码。
	if request.Method != "GET" {
		//- `request.Body`：POST/PUT 这类请求，客户端提交的数据（JSON、表单都放在 Body 里）
		//- `io.ReadAll(request.Body)`：把 Body 里面所有字节全部读出来，放到`byteDate`字节切片
		//- `string(byteDate)`：字节 → 字符串，打印到控制台
		byteDate, _ := io.ReadAll(request.Body)
		fmt.Println(string(byteDate))
	}
	fmt.Println(request.Header)
	//响应部分 writer`是发给客户端的出口
	writer.Write([]byte("Hello World!"))
}
func main() {
	//- `http.HandleFunc(路由地址, 处理函数)` 含义：**注册路由**
	//当客户端访问 `http://127.0.0.1:8080/index`，就交给上面写好的`Index`函数处理。
	http.HandleFunc("/index", Index)
	fmt.Println("http server running 127.0.0.1:8080")
	http.ListenAndServe("0.0.0.0:8080", nil)
	//这是 Go 标准库 `net/http` 启动 Web 服务的代码，属于后端，不用任何框架，原生 http 包。
	//1. `"0.0.0.0:8080"` `0.0.0.0`：监听本机**所有网卡**（本机 IP、127.0.0.1、局域网别的电脑都能访问）
	//> 如果写成 `127.0.0.1:8080`，只能你自己电脑访问，局域网其他设备连不上。
	//- `8080`：端口号。浏览器访问：`http://localhost:8080`
	//nil代表使用默认多路复用器（DefaultServeMux）
	//简单说：用系统自带的路由管理器，你可以用 `http.HandleFunc()` 注册不同 URL 对应的处理函数。

}
