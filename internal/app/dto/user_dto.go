package dto

import (
	"github.com/SalawatJoldasbaev/chat-app-golang/internal/models"
	"github.com/SalawatJoldasbaev/chat-app-golang/pkg/utility"
)

type UserRegisterDTO struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=6"`
}

func (req *UserRegisterDTO) Validate() error {
	return utility.ExtractValidationError(req)
}

func (req *UserRegisterDTO) ToModel() *models.User {
	password := utility.HashPassword(req.Password)
	return &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: password,
	}
}

type UserLoginDTO struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=6"`
}

func (req *UserLoginDTO) Validate() error {
	return utility.ExtractValidationError(req)
}

type UserUpdateDTO struct {
	Name  string `json:"name" form:"name" validate:"required"`
	Email string `json:"email" form:"email" validate:"required,email"`
}

func (req *UserUpdateDTO) ToModel() *models.User {
	return &models.User{
		Name:  req.Name,
		Email: req.Email,
	}
}

func (req *UserUpdateDTO) Validate() error {
	return utility.ExtractValidationError(req)
}

type UserChangePasswordDTO struct {
	OldPassword string `json:"old_password" form:"old_password" validate:"required,min=6"`
	NewPassword string `json:"new_password" form:"new_password" validate:"required,min=6"`
}

func (req *UserChangePasswordDTO) Validate() error {
	return utility.ExtractValidationError(req)
}
