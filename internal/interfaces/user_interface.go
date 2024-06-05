package interfaces

import (
	"github.com/SalawatJoldasbaev/chat-app-golang/internal/app/dto"
	"github.com/SalawatJoldasbaev/chat-app-golang/internal/models"
	"github.com/gofiber/fiber/v2"
)

type UserRepositoryInterface interface {
	Insert(user *models.User) (*models.User, error)
	IsEmailExists(email string, ignoreId string) bool
	FindByEmail(email string) (*models.User, error)
	FindById(id string) (*models.User, error)
	Update(id string, user *models.User) (*models.User, error)
	Delete(id string) error
	AddSession(session *models.Session) (*models.Session, error)
	FindSessionByToken(token string) (*models.Session, error)
	DeleteSession(token string) error
}

type UserServiceInterface interface {
	Register(ctx *fiber.Ctx, request *dto.UserRegisterDTO) error
	Login(ctx *fiber.Ctx, request *dto.UserLoginDTO) error
	GetMe(ctx *fiber.Ctx) error
	Logout(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx, user *dto.UserUpdateDTO) error
	ChangePassword(ctx *fiber.Ctx, user *dto.UserChangePasswordDTO) error
	Delete(ctx *fiber.Ctx) error
}

type UserHandlerInterface interface {
	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
	GetMe(ctx *fiber.Ctx) error
	Logout(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	ChangePassword(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}
