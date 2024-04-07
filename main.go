package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	if _, isSet := os.LookupEnv("PORT"); !isSet {
		os.Setenv("PORT", "8000")
	}

	addr := fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT"))
	mux := gin.Default()

	mux.GET("/", hamdleRootGet)

	mux.Run(addr)
}

func hamdleRootGet(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello world!",
	})
}
