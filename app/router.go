package app

import (
	urlShortner "go-rest-repo/app/urlShortner"

	"github.com/gin-gonic/gin"
)

func RegisterAppRoutes(r *gin.Engine) {
	urlShortner.RegisterRoutes(r)
}
