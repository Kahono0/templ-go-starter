package router

import (
	"starter/handlers"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(r *gin.Engine) {
	r.Static("/static", "./static")
	r.GET("/", handlers.IndexHandler)
}
