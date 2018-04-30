package controllers

import (
	"log"

	"github.com/aws-app/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const userKey = "username"

//handler that handles the home urls
func Home(c *gin.Context) {
	session := sessions.Default(c)
	u := session.Get(userKey)

	if u != nil {
		log.Printf("user: %v\n", u)
		user, err := models.FindUserByID(u.(uint))

		if err != nil {
			log.Println("Error getting user:", err.Error())
			c.Redirect(302, "/signup")
			return
		}

		log.Printf("Found session user: %v\n", user)
		c.Redirect(302, "/photos")
	} else {
		c.Redirect(302, "/signup")
	}
}
