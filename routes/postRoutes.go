package routes

import (
	"github.com/dayiamin/gin_blog_api/handlers"
	"github.com/dayiamin/gin_blog_api/middleware"
	"github.com/gin-gonic/gin"
)

func PostRoutes(r *gin.RouterGroup) {
	postGroup := r.Group("/post")
	{
		postGroup.GET("/", handlers.ShowPosts)
		postGroup.Use(middleware.JwtAuth())
		postGroup.POST("/register", handlers.RegisterPost)
		postGroup.DELETE("/:post_id", handlers.DeletePost)
	}
	commentGroup := postGroup.Group("/:post_id/comments")

	{
		commentGroup.POST("/", handlers.RegisterComment)
		commentGroup.DELETE("/:comment_id", handlers.DeleteComment)
	}

}
