package users

import (
	"github.com/devlopersabbir/juan_don82-server/api/users/core"
	"github.com/devlopersabbir/juan_don82-server/internal/database"
)

func Store(user *core.Users) error {
	return database.DB.Create(user).Error
}

func FindByEmail(email string) (*core.Users, error) {
	var user core.Users
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func FindAllUsers() ([]core.Users, error) {
	var users []core.Users
	err := database.DB.Find(&users).Error
	return users, err
}
