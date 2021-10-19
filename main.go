package main

import (
	"os"

	"github.com/frankenstyll/go-backend-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	router := SetupRouter()
	//r.Run(":3001") // listen and serve on 0.0.0.0:3001
	router.Run(":" + os.Getenv("GO_PORT")) //get ค่าจากไฟล์ .env ซึ่งเป็นของ go เลย

}

func SetupRouter() *gin.Engine {

	router := gin.Default()

	apiV1 := router.Group("/api/v1") //กลุ่ม url ที่เป็น version 1 จะอยู่ภายใน file ล่างนี้หมด
	routes.InitHomeRoutes(apiV1)     //routes คือชื่อ package ของ home.go และ user.go
	routes.InitUserRoutes(apiV1)     //routes คือชื่อ package ของ home.go และ user.go

	return router
}
