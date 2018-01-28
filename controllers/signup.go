package controllers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//handler that handles sign up form
func SignupForm(ctx *gin.Context) {
	session := sessions.Default(ctx)
	flashes := session.Flashes()
	session.Save()
	ctx.HTML(http.StatusOK, "signup.html", gin.H{"flash": flashes})
}
