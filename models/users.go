package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

func FindUserByID(id uint) (*User, error) {
	db, _ := setUpDB("userdata.db")
	u := &User{}

	if err := db.Where("id = ?", id).First(&u); err.Error != nil {
		return nil, err.Error
	}

	if u.ID == 0 {
		return nil, errors.New("User not found")
	}

	return u, nil
}

//gets user from database as per the username
func FindUserByUserName(userName string) (*User, error) {
	db, _ := setUpDB("userdata.db")
	u := User{}
	if err := db.Where("username = ?", userName).First(&u); err.Error != nil {
		return nil, err.Error
	}

	if u.ID == 0 {
		return nil, errors.New("User not found")
	}
	return &u, nil
}

func setUpDB(dbName string) (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", dbName)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	db.LogMode(true)
	db.AutoMigrate(&User{}, &Follower{}, &Photo{}, &Comment{})
	return db, nil
}
