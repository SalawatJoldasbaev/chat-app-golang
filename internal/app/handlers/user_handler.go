package handlers

import (
	"github.com/SalawatJoldasbaev/chat-app-golang/internal/app/dto"
	"github.com/SalawatJoldasbaev/chat-app-golang/internal/interfaces"
	"github.com/SalawatJoldasbaev/chat-app-golang/pkg/utility"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service interfaces.UserServiceInterface
}

func NewUserHandler(service interfaces.UserServiceInterface) interfaces.UserHandlerInterface {
	return &UserHandler{service: service}
}

func (u UserHandler) Register(ctx *fiber.Ctx) error {
	utility.Logger.Info("✅ USER REGISTER")
	request := new(dto.UserRegisterDTO)
	if err := ctx.BodyParser(request); err != nil {
		return utility.JsonErrorValidation(ctx, err)
	}
	return u.service.Register(ctx, request)
}

func (u UserHandler) Login(ctx *fiber.Ctx) error {
	utility.Logger.Info("✅ USER LOGIN")
	request := new(dto.UserLoginDTO)
	if err := ctx.BodyParser(request); err != nil {
		return utility.JsonErrorValidation(ctx, err)
	}
	return u.service.Login(ctx, request)
}

func (u UserHandler) GetMe(ctx *fiber.Ctx) error {
	utility.Logger.Info("✅ USER GET ME")
	return u.service.GetMe(ctx)
}

func (u UserHandler) Logout(ctx *fiber.Ctx) error {
	utility.Logger.Info("✅ USER LOGOUT")
	return u.service.Logout(ctx)
}

func (u UserHandler) Update(ctx *fiber.Ctx) error {
	utility.Logger.Info("✅ USER UPDATE")
	request := new(dto.UserUpdateDTO)
	if err := ctx.BodyParser(request); err != nil {
		return utility.JsonErrorValidation(ctx, err)
	}

	//id, err := utility.ValidateIdParams(ctx)
	//if err != nil {
	//	return err
	//}

	return u.service.Update(ctx, request)
}

func (u UserHandler) Delete(ctx *fiber.Ctx) error {
	utility.Logger.Info("✅ USER DELETE")
	return u.service.Delete(ctx)
}

func (u UserHandler) ChangePassword(ctx *fiber.Ctx) error {
	utility.Logger.Info("✅ USER CHANGE PASSWORD")
	request := new(dto.UserChangePasswordDTO)
	if err := ctx.BodyParser(request); err != nil {
		return utility.JsonErrorValidation(ctx, err)
	}
	return u.service.ChangePassword(ctx, request)
}
