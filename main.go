package main

import (
	"log"
	"os"

	"github.com/aws-app/controllers"
	"github.com/aws-app/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	logging "github.com/sirupsen/logrus"
)

func main() {
	db, err := gorm.Open("sqlite3", "userdata.db")
	if err != nil {
		logging.Fatalf("Unable to connect to database, got error %s", err)
	}
	defer db.Close()
	db.LogMode(true)
	db.AutoMigrate(&models.User{}, &models.Follower{}, &models.Photo{}, &models.Comment{})

	r := controllers.StartUp()
	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	log.Printf("Listening on port %s\n", port)
	r.Run(":" + port)
}
