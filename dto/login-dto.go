package dto

type LoginDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponseDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type RegisterDTO struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
