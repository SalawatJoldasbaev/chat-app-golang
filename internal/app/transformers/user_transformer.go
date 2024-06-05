package transformers

import (
	"github.com/SalawatJoldasbaev/chat-app-golang/internal/models"
	"github.com/SalawatJoldasbaev/chat-app-golang/pkg/constants"
)

type UserLoginResponse struct {
	Token string `json:"token"`
}

func UserLoginTransformer(token string) *UserLoginResponse {
	return &UserLoginResponse{Token: token}
}

type UserResponse struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	RegisteredAt string `json:"registered_at"`
}

func UserTransformer(user *models.User) *UserResponse {
	return &UserResponse{
		ID:           user.Id.String(),
		Name:         user.Name,
		Email:        user.Email,
		RegisteredAt: user.CreatedAt.Format(constants.TimestampFormat),
	}
}
