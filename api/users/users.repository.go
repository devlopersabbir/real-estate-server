package users

import (
	"github.com/devlopersabbir/juan_don82-server/api/users/core"
	"github.com/devlopersabbir/juan_don82-server/internal/database"
)

func Store(user *core.Users) error {
	return database.DB.Create(user).Error
}
