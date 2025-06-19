package main

import (
	"go-rest-repo/app"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	app.RegisterAppRoutes(r)
	log.Println("Server running on :8080")
	r.Run(":8080")
}
