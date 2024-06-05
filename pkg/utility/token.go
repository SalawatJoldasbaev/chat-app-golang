package utility

import (
	"github.com/SalawatJoldasbaev/chat-app-golang/configs"
	"github.com/SalawatJoldasbaev/chat-app-golang/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"strings"
	"time"
)

type JwtCustomClaims struct {
	Issuer string `json:"issuer"`
	jwt.StandardClaims
}
type Token struct {
	Token     string
	ExpiresAt time.Time
}

func TokenGenerate(user *models.User) (*Token, error) {
	expireHour, _ := time.ParseDuration("48h")
	expiresAt := time.Now().Add(time.Hour + expireHour)

	claims := JwtCustomClaims{
		Issuer: user.Id.String(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt.Unix(),
		},
	}
	jwtSecret := configs.Configs.Jwt.Secret
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return nil, err
	}
	return &Token{
		Token:     tokenString,
		ExpiresAt: expiresAt,
	}, nil
}

func GetAuthorizationToken(ctx *fiber.Ctx) string {
	authorizationToken := string(ctx.Request().Header.Peek("Authorization"))
	return strings.Replace(authorizationToken, "Bearer ", "", 1)
}
