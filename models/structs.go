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
func NewUser(email,username,password, fullName string) (*User){
	return &User{
		Email: email,
		Username: username,
		Password: password,
		FullName:fullName,
	}
}

func NewFollower(userID, followerId uint) *Follower{
	return &Follower{
		UserID: userID,
		FollowerID: followerId,
	}
}

func NewPhoto(userId,likes uint, fileName,caption string, createdAt time.Time) *Photo{
	return &Photo{
		UserID:userId,
		Filename: fileName,
		Caption: caption,
		Likes: likes,
		CreatedAt: createdAt,
	}
}
func NewComment(userID, photoID uint, text string, createdAt time.Time) *Comment{
	return &Comment{
		UserID: userID,
		PhotoID: photoID,
		Text: text,
		CreatedAt: createdAt,
	}
}