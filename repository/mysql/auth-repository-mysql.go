package mysql

import (
	"github.com/devnura/go-echo-rest-api/entity"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) FindByEmail(email string) (*entity.User, error) {

	var user entity.User

	statement := r.db.Table("users").Where("email = ?", email).First(&user)

	if err := statement.Error; err != nil {
		return nil, err
	}

	return &user, nil
}
