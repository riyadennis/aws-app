package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerRoutes() *gin.Engine {
	r := gin.Default()
	r.NoRoute(noRoute)
	r.LoadHTMLGlob("templates/**/*.html")
	r.Static("/public", "./public")
	return r
}

func noRoute(ctx *gin.Context) {
	ctx.HTML(http.StatusNotFound, "404.html", nil)
}
