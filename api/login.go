package api

import (
	"github.com/gin-gonic/gin"
	"myfeed/model"
	"net/http"
)

func Login(c *gin.Context) {
	var formData model.User
	_ = c.ShouldBindJSON(&formData)
	//var code int

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
	})

}
