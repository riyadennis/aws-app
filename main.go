package main

import (
	"os"

	"github.com/aws-app/controllers"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	logging "github.com/sirupsen/logrus"
)

func main() {
	r := controllers.StartUp()
	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	logging.Printf("Listening on port %s\n", port)
	r.Run(":" + port)
}
