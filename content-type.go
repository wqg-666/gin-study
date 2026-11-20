package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		// 1. 获取 Content-Type
		contentType := r.Header.Get("Content-Type")
		fmt.Printf("\n--- 收到请求，Content-Type 为: %s ---\n", contentType)

		// 2. 根据 Content-Type 分支处理
		if contentType == "application/json" {
			// 解析 JSON
			var data map[string]interface{}
			body, _ := io.ReadAll(r.Body) // 先读出原始内容看看
			fmt.Printf("原始 JSON 内容: %s\n", string(body))
			json.Unmarshal(body, &data) // 解析
			fmt.Printf("解析后的 Map: %+v\n", data)

		} else if contentType == "application/x-www-form-urlencoded" {
			// 解析 x-www-form-urlencoded
			r.ParseForm() // 官方库必须调用这一步
			fmt.Printf("解析后的表单: %+v\n", r.PostForm)
			fmt.Printf("获取单个字段 name: %s\n", r.PostFormValue("name"))

		} else if contentType == "multipart/form-data" {
			// 解析 form-data
			// 10 << 20 表示允许 10MB 的文件读入内存，超过的存到临时文件
			err := r.ParseMultipartForm(10 << 20)
			if err != nil {
				fmt.Println("解析 form-data 失败:", err)
				return
			}
			fmt.Printf("文本字段: %+v\n", r.MultipartForm.Value)
			// 尝试获取文件
			file, handler, err := r.FormFile("myfile")
			if err == nil {
				defer file.Close()
				fmt.Printf("获取到文件: %s, 大小: %d bytes\n", handler.Filename, handler.Size)
			}
		} else {
			fmt.Println("未知的 Content-Type，直接读取 Body:")
			body, _ := io.ReadAll(r.Body)
			fmt.Println(string(body))
		}
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
