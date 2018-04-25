package models

import (
	"time"
)

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
