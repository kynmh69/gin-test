package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kynmh69/go-mysql/database"
)

func init() {
	database.ConnectToMySQL()
}

func main() {
	r := gin.New()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
