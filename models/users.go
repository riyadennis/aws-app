package models

import (
	"errors"
)

func FindUserByID(id uint) (*User, error) {
	db := setUpDB()
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
	db := setUpDB()
	u := User{}
	if err := db.Where("username = ?", userName).First(&u); err.Error != nil {
		return nil, err.Error
	}

	if u.ID == 0 {
		return nil, errors.New("User not found")
	}
	return &u, nil
}
