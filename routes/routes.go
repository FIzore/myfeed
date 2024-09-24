package routes

import "github.com/gin-gonic/gin"

func login(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}
func register(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}
func reset(c *gin.Context) {
	c.HTML(200, "reset.html", nil)
}

//	r.GET("/", func(c *gin.Context) {
//		c.HTML(200, "basic.html", nil)
//
// })
func report(c *gin.Context) {
	c.HTML(200, "report.html", nil)
}
func manual(c *gin.Context) {
	c.HTML(200, "manual.html", nil)
}
func privacy(c *gin.Context) {
	c.HTML(200, "privacy.html", nil)
}
