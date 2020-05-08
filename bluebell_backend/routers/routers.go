package routers

import (
	"bluebell_backend/controller"
	"bluebell_backend/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.POST("/login", controller.LoginHandler)
	r.POST("/register", controller.RegisterHandler)
	v1 := r.Group("/api/v1", controller.JWTAuthMiddleware())
	{
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDetailHandler)

		v1.POST("/post", controller.CreatePostHandler)
		v1.GET("/post/:id", controller.PostDetailHandler)
		v1.GET("/post", controller.PostListHandler)

		v1.POST("/vote", controller.VoteHandler)

		v1.POST("/comment", controller.CommentHandler)
		v1.GET("/comment", controller.CommentListHandler)

	}

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}