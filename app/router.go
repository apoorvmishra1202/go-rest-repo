package app

import (
	urlshortner "go-rest-repo/app/urlShortner"

	"github.com/gin-gonic/gin"
)

func RegisterAppRoutes(r *gin.Engine) {
	urlshortner.RegisterRoutes(r)
}
