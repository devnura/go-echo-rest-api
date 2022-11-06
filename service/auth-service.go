package service

import (
	"errors"
	"fmt"

	"github.com/devnura/go-echo-rest-api/common/log"
	"github.com/devnura/go-echo-rest-api/dto"
	"github.com/devnura/go-echo-rest-api/entity"
	"github.com/devnura/go-echo-rest-api/repository"
	"github.com/devnura/go-echo-rest-api/transfer"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AuthService interface {
	VerifyCredential(c echo.Context, req *dto.LoginDTO) (*entity.User, error)
}

type authService struct {
	authRepository repository.AuthRepository
}

func NewAuthService(u repository.AuthRepository) *authService {
	return &authService{
		authRepository: u,
	}
}

func (s *authService) VerifyCredential(c echo.Context, req *dto.LoginDTO) (*entity.User, error) {

	fmt.Println(req)
	var entity entity.User
	err := copier.Copy(&entity, req)
	if err != nil {
		log.InfoWithID(c, "step 1.1")
		return nil, err
	}

	user, err := s.authRepository.FindByUsername(&entity)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, transfer.NewCustomError("Invalid email or password")
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}
