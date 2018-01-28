package models

import "errors"

func FindUserByID(id uint) (*User, error) {
	u := &User{}

	if err := db.Where("id = ?", id).First(&u); err.Error != nil {
		return nil, err.Error
	}

	if u.ID == 0 {
		return nil, errors.New("User not found")
	}

	return u, nil
}
