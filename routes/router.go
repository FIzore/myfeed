package routes

import "github.com/gin-gonic/gin"

func Login(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}
func Register(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}
func Reset(c *gin.Context) {
	c.HTML(200, "reset.html", nil)
}

func Index(c *gin.Context) {
	c.HTML(200, "index.html", nil)

}
func Report(c *gin.Context) {
	c.HTML(200, "report.html", nil)
}
func Manual(c *gin.Context) {
	c.HTML(200, "manual.html", nil)
}
func Privacy(c *gin.Context) {
	c.HTML(200, "privacy.html", nil)
}
