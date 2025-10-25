package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./login.html", "./index.html")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	r.POST("/login", func(c *gin.Context) {
		// key根据input的name来的
		username := c.PostForm("username")
		password := c.PostForm("password")
		//username := c.DefaultPostForm("username", "admin")
		//password := c.DefaultPostForm("xxx", "123456")

		//username, ok := c.GetPostForm("username")
		//if !ok {
		//	username = "admin"
		//}
		//
		//password, ok := c.GetPostForm("xxx")
		//if !ok {
		//	password = "123456"
		//}
		c.HTML(http.StatusOK, "index.html", gin.H{
			"username": username,
			"password": password,
		})
	})

	r.Run(":8080")
}
