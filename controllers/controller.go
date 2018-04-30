package controllers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

// Cognito represents an abstraction of the Amazon Cognito identity provider service.
type Cognito struct {
	cip *cognitoidentityprovider.CognitoIdentityProvider
}

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

	photos := r.Group("/photos", AuthRequired())
	{
		photos.POST("/", CreatePhoto)
		//photos.GET("/", FetchAllPhotos)
		//photos.GET("/:id", FetchSinglePhoto)
		//photos.PUT("/:id", UpdatePhoto)
		//photos.DELETE("/:id", DeletePhoto)
		//photos.POST("/:id/like", LikePhoto)
		//photos.POST("/:id/comment", CommentPhoto)
	}
	return r
}

func notFoundHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusNotFound, "404.html", nil)
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// NewCognito creates a new instance of the Cognito client
func NewCognito() *Cognito {

	c := &Cognito{}

	// Create Session
	sess := session.Must(session.NewSession())
	c.cip = cognitoidentityprovider.New(sess)

	return c
}
