package main

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("users", func(c *gin.Context) {
		//读取 form-data 表单中 key 为`file`的单个上传文件
		fileHeader, err := c.FormFile("file")
		//`fileHeader`，它里面保存了本次上传文件的全部元信息
		if err != nil {
			//err 代表：Gin 从本次请求的 form-data 表单里，拿不到 key=file 的这个文件信息。
			//从 HTTP 请求中提取文件信息失败
			fmt.Println(err)
			return
		}
		fmt.Println(fileHeader.Filename) //文件名
		fmt.Println(fileHeader.Size)     //文件大小，单位是字节
		//方法一
		file, _ := fileHeader.Open()                    //打开上传得到的文件流
		byteDate, _ := io.ReadAll(file)                 //一次性读取文件全部内容到 byteDate 字节切片
		os.WriteFile("uploads/xxx.jpg", byteDate, 0666) // 将字节数据写入本地磁盘，保存为 xxx.jpg
		fmt.Println(err)
		//方法二
		//`uploads/xxx/yyy/`：存放目录 ,`fileHeader.Filename`：前端上传的原始文件名
		//err1 := c.SaveUploadedFile(fileHeader, "uploads/xxx/yyy/"+fileHeader.Filename)
		//err1：把文件写到服务器硬盘失败
		//fmt.Println(err1)
		//方法一：一次性把整个文件全部读到内存，再一次性写入磁盘；大文件容易内存溢出 OOM。
		//方法二：流式读写，边读上传流、边写磁盘，不会一次性加载全部文件到内存，工程推荐。
	})
	r.Run(":8080")
}
