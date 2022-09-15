package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHello(ctx *gin.Context) {
	ctx.Data(200, "text/plain", []byte("Hello, It Home!"))
}

func DeleteHello(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.String(http.StatusOK, "hello DELETE %s", id)
}