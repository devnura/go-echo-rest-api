package repository

import (
	"github.com/devnura/go-echo-rest-api/entity"
	"gorm.io/gorm"
)

type AuthRepository interface {
	FindByUsername(user *entity.User) (*entity.User, error)
}

type authConnection struct {
	connection *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authConnection{connection: db}
}

func (db *authConnection) FindByUsername(user *entity.User) (*entity.User, error) {

	statement := db.connection.Where("email = ?", user.Email).First(&user)

	if err := statement.Error; err != nil {
		return nil, err
	}

	return user, nil

}
