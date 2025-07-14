package routes

import (
	"github.com/dayiamin/gin_blog_api/handlers"
	"github.com/dayiamin/gin_blog_api/middleware"
	"github.com/gin-gonic/gin"
)



func UserRoutes(r *gin.RouterGroup){
	userGroup := r.Group("/user")
	
	userGroup.POST("/register",handlers.RegisterUser)
	userGroup.POST("/login",handlers.Login)
	// userGroup.POST("/profile",handlers.CreateProfile).Use(middleware.JwtAuth())
	profileGroup := userGroup.Group("/profile")
	profileGroup.GET("/:user_name",handlers.ShowProfile)
	profileGroup.Use(middleware.JwtAuth())
	{
		profileGroup.POST("/", handlers.CreateProfile)
	}

}