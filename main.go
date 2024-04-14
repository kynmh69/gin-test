package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/graceful"
	"github.com/gin-gonic/gin"
	"github.com/kynmh69/go-mysql/database"
)

func init() {
	database.ConnectToMySQL()
}

type Msg struct {
	Message string `json:"message"`
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost"}

	r, err := graceful.Default()
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Close()
	// r.Use(logger.SetLogger())
	r.Use(cors.New(config))

	r.GET("/ping", func(ctx *gin.Context) {
		msg := Msg{Message: "pong"}
		ctx.JSON(http.StatusOK, msg)
	})

	if err := r.RunWithContext(ctx); err != nil && err != context.Canceled {
		log.Fatalln(err)
	}
}
