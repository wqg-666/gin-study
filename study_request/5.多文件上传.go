package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("users", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			fmt.Println(err)
			return
		}

		//遍历表单里所有form-data分组，再遍历每个分组里的每一个上传文件，
		for _, headers := range form.File { //遍历 `form.File` 这个 map。
			for _, header := range headers { //`headers`是一组文件切片，这层循环遍历组内每一个单独文件
				c.SaveUploadedFile(header, "uploads/"+header.Filename)
			}
		}
	})
	r.Run(":8080")
}
