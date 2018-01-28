package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

//struct that will define user table
type User struct {
	ID       uint
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
	ID        uint
	UserID    uint
	Filename  string
	Caption   string
	CreatedAt time.Time
	Likes     uint
}
type Comment struct {
	ID        uint
	UserID    uint
	PhotoID   uint
	Text      string
	CreatedAt time.Time
}
