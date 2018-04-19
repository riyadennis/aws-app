package lib

import (
	"github.com/aws-app/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func SetUpDB(dbName string) error {
	db, err := gorm.Open("sqlite3", dbName)
	if err != nil {
		return err
	}
	defer db.Close()
	db.LogMode(true)
	db.AutoMigrate(&models.User{}, &models.Follower{}, &models.Photo{}, &models.Comment{})
	return nil
}
