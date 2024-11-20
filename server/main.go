package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"server/src/config"
	v1 "server/src/controllers/v1"
	_ "server/src/docs"
	"server/src/middleware"
	"server/src/router"
	"server/src/service/flow"
)

// @title AREA - API
// @version 2.0
// @description This API is for the project AREA (A website like IFTT or Zappier)
// @BasePath /
// @securityDefinitions.apikey Session Token
// @in header
// @name Authorization
func main() {
	fmt.Println("Starting server..")
	config.InitDB()
	config.InitRedis()

	go flow.RunFlow()

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Use(middleware.CORSMiddleware())
	r.GET("/about.json", v1.About)
	v1Router := r.Group("/v1")
	router.SetupV1Router(v1Router)

	err := r.Run("0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
}
