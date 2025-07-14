package main

import (

	"github.com/dayiamin/gin_blog_api/database"
	"github.com/dayiamin/gin_blog_api/routes"
	_ "github.com/dayiamin/gin_blog_api/docs"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
)

// TODO  write handler for creating user login users and creating posts and comments

// load the init file
func init(){
	var envErr = godotenv.Load()
	if envErr != nil{
		log.Fatal("env error")
	}
}


// @title Blog Post api
// @version 1.0
// @description this is api for creating users and posting blogs and comments
// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.apikey JWT
// @in header
// @name Authorization
func main(){

	router := gin.Default()
	v1Router := router.Group("/api/v1")
	v1Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	db.Connect()
	
	routes.UserRoutes(v1Router)
	routes.PostRoutes(v1Router)

	router.Run(":8080")			
}