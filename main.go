package main

import (
	"__module_name__/internal"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.New()
	app.GET("/healthz", internal.Healthz)
	app.Run()
}
