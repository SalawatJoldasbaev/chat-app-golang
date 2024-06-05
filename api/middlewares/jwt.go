package middlewares

import (
	"errors"
	"github.com/SalawatJoldasbaev/chat-app-golang/configs"
	"github.com/SalawatJoldasbaev/chat-app-golang/pkg/utility"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type JwtCustomClaims struct {
	Issuer string `json:"issuer"`
	jwt.StandardClaims
}

type SkipperRoutesData struct {
	Method  string
	UrlPath string
}

func JwtMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		ctx.Set("X-XSS-Protection", "1; mode=block")
		ctx.Set("Strict-Transport-Security", "max-age=5184000")
		ctx.Set("X-DNS-Prefetch-Control", "off")

		// skip whitelist routes
		for _, whiteList := range whiteListRoutes() {
			if ctx.Method() == whiteList.Method && ctx.Path() == whiteList.UrlPath {
				return ctx.Next()
			}
		}
		// check header token
		authorizationToken := utility.GetAuthorizationToken(ctx)
		if authorizationToken == "" {
			err := errors.New("missing Bearer token")
			return utility.JsonErrorUnauthorized(ctx, err)
		}
		// verify token
		jwtToken, err := jwt.ParseWithClaims(authorizationToken, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(configs.Configs.Jwt.Secret), nil
		})

		if err != nil {
			return utility.JsonErrorUnauthorized(ctx, err)
		}

		claimsData := jwtToken.Claims.(*JwtCustomClaims)
		utility.Logger.Info("✅ SET USER AUTH")
		ctx.Locals("user_auth", claimsData.Issuer)
		return ctx.Next()
	}
}

func whiteListRoutes() []SkipperRoutesData {
	return []SkipperRoutesData{
		{"POST", "/api/auth/register"},
		{"POST", "/api/auth/login"},
	}
}
