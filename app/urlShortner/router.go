package urlshortner

import (
	"go-rest-repo/app/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	group := r.Group("/urlshortner", middlewares.Auth())
	{
		group.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "pong from urlShortner"})
		})
		group.POST("/shorten", CreateShortURL)
		group.GET("/:short", GetOriginalURLHandler)
	}
}
