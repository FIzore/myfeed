package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"myfeed/routes/login"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("template/*")
	r.GET("/login", login.Login)
	r.GET("/register", func(c *gin.Context) {
		c.HTML(200, "register.html", nil)
	})
	r.GET("/password/reset", func(c *gin.Context) {
		c.HTML(200, "reset.html", nil)
	})
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "basic.html", nil)

	})
	r.GET("/report", func(c *gin.Context) {
		c.HTML(200, "report.html", nil)
	})
	r.GET("/manual", func(c *gin.Context) {
		c.HTML(200, "manual.html", nil)
	})
	r.GET("/privacy", func(c *gin.Context) {
		c.HTML(200, "privacy.html", nil)
	})
	r.Run(":8080")
}
