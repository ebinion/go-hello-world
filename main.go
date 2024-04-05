package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	mux := gin.Default()

	mux.GET("/", hamdleRootGet)

	mux.Run()
}

func hamdleRootGet(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello world!",
	})
}
