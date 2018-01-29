package controllers

import (
	"log"
	"net/http"

	"github.com/aws-app/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

//handler that handles sign up form
func SignupForm(ctx *gin.Context) {
	session := sessions.Default(ctx)
	flashes := session.Flashes()
	session.Save()
	ctx.HTML(http.StatusOK, "signup.html", gin.H{"flash": flashes})
}
func Signup(ctx *gin.Context) {
	password := []byte(ctx.PostForm("password"))
	hashedPassword, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	user := &models.User{
		FullName: ctx.PostForm("fullName"),
		Username: ctx.PostForm("username"),
		Email:    ctx.PostForm("email"),
		Password: string(hashedPassword),
	}
	session := sessions.Default(ctx)
	if !validateUserName(ctx, session, user) {
		log.Println("Creating DB user:", user.Username)
	}

}

func validateUserName(ctx *gin.Context, session sessions.Session, user *models.User) bool {
	u, _ := models.FindUserByUserName(ctx.PostForm("username"))
	if u != nil {
		message := "This username is not available, please try another one"
		session.AddFlash(message)
		ctx.HTML(
			http.StatusOK,
			"signup.html",
			gin.H{
				"flash": session.Flashes(),
				"user":  user,
			},
		)
		return false
	}
	return true
}
