package service

import (
	"log"

	"github.com/devnura/go-echo-rest-api/dto"
	"github.com/devnura/go-echo-rest-api/entity"
	"github.com/devnura/go-echo-rest-api/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(loginDTO *dto.LoginDTO) interface{}
}

type authService struct {
	authRepository repository.AuthRepository
}

func NewAuthService(authRepo repository.AuthRepository) AuthService {
	return &authService{
		authRepository: authRepo,
	}
}

func (service *authService) Login(dto *dto.LoginDTO) interface{} {
	res := service.authRepository.FindByUsername(dto.Email)
	if v, ok := res.(entity.User); ok {
		comparedPassword := comparePassword(v.PasswordHash, []byte(dto.Password))
		if v.Email == dto.Email && comparedPassword {
			return res
		}
		return false
	}
	return false
}

func comparePassword(hasherPassword string, password []byte) bool {
	byteHash := []byte(hasherPassword)
	err := bcrypt.CompareHashAndPassword(byteHash, password)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
