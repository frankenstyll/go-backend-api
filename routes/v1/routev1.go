package v1

import "github.com/gin-gonic/gin"

func InitRoutesV1(rg *gin.RouterGroup) {

	routerGroup := rg.Group("/")
	routerGroup.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"API_VERSION": "1.0.0",
			"message":     "home",
		})
	})
}
