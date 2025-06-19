package urlshortner

import (
	"go-rest-repo/internal/network"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateShortURL(c *gin.Context) {
	var req ShortenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		network.SendFailure(c, http.StatusBadRequest, err.Error())
		return
	}
	short := SetShortURL(req.URL)
	network.SendSuccess(c, ShortenResponse{ShortURL: short})
}

func GetOriginalURLHandler(c *gin.Context) {
	short := c.Param("short")
	if orig, ok := GetOriginalURL(short); ok {
		network.SendSuccess(c, gin.H{"url": orig})
		return
	}
	network.SendFailure(c, http.StatusNotFound, "Short URL not found")
}
