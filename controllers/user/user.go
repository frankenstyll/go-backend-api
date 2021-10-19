package usercontroller

import "github.com/gin-gonic/gin"

func GetAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "user",
	})
}

func Register(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "register",
	})
}
