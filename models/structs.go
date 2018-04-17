package models

import (
	"time"

	"github.com/jinzhu/gorm"
	logging "github.com/sirupsen/logrus"
)

var db *gorm.DB

//struct that will define user table
type User struct {
	ID       uint `gorm:"AUTO_INCREMENT"`
	Email    string
	Username string
	Password string
	FullName string
}

type Follower struct {
	UserID     uint `gorm:"unique_index:idx_user_follower"`
	FollowerID uint `gorm:"unique_index:idx_user_follower"`
}

type Photo struct {
	ID        uint `gorm:"AUTO_INCREMENT"`
	UserID    uint
	Filename  string
	Caption   string
	CreatedAt time.Time
	Likes     uint
}
type Comment struct {
	ID        uint `gorm:"AUTO_INCREMENT"`
	UserID    uint
	PhotoID   uint
	Text      string
	CreatedAt time.Time
}

func setUpDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "userdata.db")
	if err != nil {
		logging.Fatalf("Unable to connect to database, got error %s", err)
	}
	defer db.Close()
	db.LogMode(true)
	db.AutoMigrate(&User{}, &Follower{}, &Photo{}, &Comment{})
	return db
}
