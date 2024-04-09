package repository

import (
	"errors"

	"app/database"
	"app/model"

	"gorm.io/gorm"
)

func GetUserByEmail(e string) (*model.User, error) {
	db := database.DB
	var user model.User
	err := db.Where(&model.User{Email: e}).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
