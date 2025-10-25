package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

// 1. 访问http://127.0.0.1:8080/index, 上传文件
// 2. 点击上传后, 会出发POST请求, 将上传的文件保存到服务器本地

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("index.html")
	r.GET("/index", func(c *gin.Context) {

		c.HTML(http.StatusOK, "index.html", nil)
	})

	// 单个文件上传
	r.POST("/upload", func(c *gin.Context) {
		// 从请求中读取文件
		f, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			// 将读取到的文件保存在本地 (服务端)
			//dst := fmt.Sprintf("./%s", f.Filename)
			dst := path.Join("./", f.Filename)
			c.SaveUploadedFile(f, dst)
			c.JSON(http.StatusOK, gin.H{
				"message": fmt.Sprintf("'%s' uploaded!", f.Filename),
			})
		}
	})

	r.Run(":8080")
}
