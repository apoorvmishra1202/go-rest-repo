package network

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func SendFailure(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"error": message})
}
