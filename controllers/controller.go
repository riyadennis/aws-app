package controllers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func StartUp() *gin.Engine {
	r := gin.Default()
	store := sessions.NewCookieStore([]byte("viErkShjgQP59tgelRXsILXNEarwRA6p"))
	r.Use(sessions.Sessions("photos-session", store))
	r.NoRoute(notFoundHandler)
	r.LoadHTMLGlob("templates/**/*.html")
	r.Static("/public", "./public")
	r.GET("/", Home)
	r.GET("/signup", SignupForm)
	r.POST("/signup", Signup)
	return r
}

func notFoundHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusNotFound, "404.html", nil)
}
