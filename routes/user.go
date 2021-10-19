package routes

import (
	usercontroller "github.com/frankenstyll/go-backend-api/controllers/user"
	"github.com/gin-gonic/gin"
)

func InitUserRoutes(rg *gin.RouterGroup) {

	routerGroup := rg.Group("/users") // localhost:3000/api/v1/users

	// localhost:3000/api/v1/users
	routerGroup.GET("/", usercontroller.GetAll) //เรียก

	// localhost:3000/api/v1/users/register
	routerGroup.POST("/register", usercontroller.Register)

}
