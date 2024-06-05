package utility

import (
	"errors"
	"github.com/SalawatJoldasbaev/chat-app-golang/pkg/constants"
	"reflect"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ExtractValidationError(req interface{}) error {
	var message string
	var v = validator.New()

	err := v.RegisterValidation("date", validateDate)
	// get json tag
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}

		return name
	})

	err = v.Struct(req)
	if err != nil {
		for i, err := range err.(validator.ValidationErrors) {
			if i > 0 {
				message += " | "
			}

			message += err.Field() + ": " + err.Tag()
		}

		return errors.New(message)
	}

	return nil
}

func ValidateIdParams(ctx *fiber.Ctx) (string, error) {
	id := ctx.Params("id")
	if id == "" {
		return "", JsonErrorValidation(ctx, constants.ErrIdRequired)
	}

	return id, nil
}

func validateDate(fl validator.FieldLevel) bool {
	timeStr := fl.Field().String()
	_, err := time.Parse(constants.DateFormat, timeStr)
	return err == nil
}
