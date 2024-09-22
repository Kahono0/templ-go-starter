package handlers

import (
	"starter/views"


	"github.com/gin-gonic/gin"
)

func IndexHandler(ctx *gin.Context) {
	t := views.Index()

	render(ctx, 200, t)
}
