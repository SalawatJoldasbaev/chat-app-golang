package services

import (
	"errors"
	"github.com/SalawatJoldasbaev/chat-app-golang/internal/app/dto"
	"github.com/SalawatJoldasbaev/chat-app-golang/internal/app/transformers"
	"github.com/SalawatJoldasbaev/chat-app-golang/internal/interfaces"
	"github.com/SalawatJoldasbaev/chat-app-golang/internal/models"
	"github.com/SalawatJoldasbaev/chat-app-golang/pkg/utility"
	"github.com/gofiber/fiber/v2"
)

type UserService struct {
	repository interfaces.UserRepositoryInterface
}

func NewUserService(repository interfaces.UserRepositoryInterface) interfaces.UserServiceInterface {
	return &UserService{repository: repository}
}

func (u *UserService) Register(ctx *fiber.Ctx, request *dto.UserRegisterDTO) error {
	if err := request.Validate(); err != nil {
		return utility.JsonErrorValidation(ctx, err)
	}
	if u.repository.IsEmailExists(request.Email, "") {
		err := errors.New("email already exists")
		return utility.JsonError(ctx, err, "E_EMAIL_EXISTS")
	}
	user := request.ToModel()
	_, err := u.repository.Insert(user)
	if err != nil {
		return utility.JsonErrorInternal(ctx, err, "E_USER_ADD")
	}
	token, err := utility.TokenGenerate(user)
	if err != nil {
		return utility.JsonErrorInternal(ctx, err, "E_USER_TOKEN")
	}
	_, err = u.repository.AddSession(&models.Session{
		UserId:  user.Id,
		Data:    token.Token,
		Expires: token.ExpiresAt,
	})
	if err != nil {
		return utility.JsonErrorInternal(ctx, err, "E_USER_ADD")
	}
	return utility.JsonSuccess(ctx, transformers.UserLoginTransformer(token.Token))
}

func (u *UserService) Login(ctx *fiber.Ctx, request *dto.UserLoginDTO) error {
	if err := request.Validate(); err != nil {
		return utility.JsonErrorValidation(ctx, err)
	}
	user, err := u.repository.FindByEmail(request.Email)
	if err != nil {
		return utility.JsonErrorUnauthorized(ctx, errors.New("username or password is wrong"))
	}
	if !utility.ComparePassword(user.Password, request.Password) {
		return utility.JsonErrorUnauthorized(ctx, errors.New("username or password is wrong"))
	}
	token, err := utility.TokenGenerate(user)
	if err != nil {
		return utility.JsonErrorInternal(ctx, err, "E_USER_TOKEN")
	}
	return utility.JsonSuccess(ctx, transformers.UserLoginTransformer(token.Token))
}

func (u *UserService) GetMe(ctx *fiber.Ctx) error {
	id := ctx.Locals("user_auth").(string)
	user, err := u.repository.FindById(id)
	if err != nil {
		return utility.JsonErrorInternal(ctx, err, "E_USER_GET")
	}
	return utility.JsonSuccess(ctx, transformers.UserTransformer(user))
}

func (u *UserService) Logout(ctx *fiber.Ctx) error {
	token := utility.GetAuthorizationToken(ctx)
	err := u.repository.DeleteSession(token)
	if err != nil {
		return utility.JsonErrorInternal(ctx, err, "E_USER_LOGOUT")
	}
	return utility.JsonSuccess(ctx, fiber.Map{"token": token})
}

func (u *UserService) Update(ctx *fiber.Ctx, user *dto.UserUpdateDTO) error {
	if err := user.Validate(); err != nil {
		return utility.JsonErrorValidation(ctx, err)
	}
	id := ctx.Locals("user_auth").(string)
	userModel, err := u.repository.FindById(id)
	if err != nil {
		return utility.JsonErrorNotFound(ctx, err)
	}
	if user.Email != userModel.Email && u.repository.IsEmailExists(user.Email, id) {
		err := errors.New("email already exists")
		return utility.JsonError(ctx, err, "E_EMAIL_EXISTS")
	}
	newUserData, err := u.repository.Update(id, &models.User{
		Id:        userModel.Id,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: userModel.CreatedAt,
	})
	if err != nil {
		return utility.JsonErrorInternal(ctx, err, "E_USER_UPDATE")
	}
	return utility.JsonSuccess(ctx, transformers.UserTransformer(newUserData))
}

func (u *UserService) ChangePassword(ctx *fiber.Ctx, user *dto.UserChangePasswordDTO) error {
	if err := user.Validate(); err != nil {
		return utility.JsonErrorValidation(ctx, err)
	}
	id := ctx.Locals("user_auth").(string)
	userModel, err := u.repository.FindById(id)
	if err != nil {
		return utility.JsonErrorNotFound(ctx, err)
	}
	if !utility.ComparePassword(userModel.Password, user.OldPassword) {
		err := errors.New("old password is wrong")
		return utility.JsonError(ctx, err, "E_PASSWORD_WRONG")
	}
	userModel.Password = utility.HashPassword(user.NewPassword)
	_, err = u.repository.Update(id, userModel)
	if err != nil {
		return utility.JsonErrorInternal(ctx, err, "E_USER_UPDATE")
	}
	return utility.JsonSuccess(ctx, fiber.Map{"id": userModel.Id.String()})
}

func (u *UserService) Delete(ctx *fiber.Ctx) error {
	id := ctx.Locals("user_auth").(string)
	err := u.repository.Delete(id)
	if err != nil {
		return utility.JsonErrorInternal(ctx, err, "E_USER_DELETE")
	}
	return utility.JsonSuccess(ctx, fiber.Map{"id": id})
}
