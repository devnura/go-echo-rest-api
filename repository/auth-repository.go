package repository

import (
	"log"

	"github.com/devnura/go-echo-rest-api/entity"
	"gorm.io/gorm"
)

type AuthRepository interface {
	FindByUsername(email string) interface{}
}

type authConnection struct {
	connection *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authConnection{
		connection: db,
	}
}

func (db *authConnection) FindByUsername(email string) interface{} {
	var user entity.User
	res := db.connection.Where("email = ?", email).Take(&user)
	if res.Error != nil {
		log.Printf("%v", res.Error)
		return nil
	}
	return user
}
