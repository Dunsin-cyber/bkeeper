package middlewares

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Dunsin-cyber/bkeeper/common"
	"github.com/Dunsin-cyber/bkeeper/internal/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AppMiddleware struct {
	Logger echo.Logger
	DB     *gorm.DB
}

func (appMiddleware *AppMiddleware) AuthenticationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {
		// For example, check for a valid JWT token in the request header
		//supply jwt
		c.Response().Header().Add("Vary", "Authorization")
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer") {
			errMsg := "Missing or invalid Authorization header"
			return common.SendUnauthorizedResponse(c, &errMsg)
		}
		authHeaderSpit := strings.Split(authHeader, " ")
		if len(authHeaderSpit) != 2 {
			errMsg := "Invalid Authorization header format"
			return common.SendUnauthorizedResponse(c, &errMsg)
		}
		tokenString := authHeaderSpit[1]

		claims, err := common.ParseJWTSignedAccessToken(tokenString)
		if err != nil {
			fmt.Println(err)
			errMsg := err.Error()
			return common.SendUnauthorizedResponse(c, &errMsg)
		}

		if common.IsClaimExpired(claims) {
			errMsg := "Token has expired"
			return common.SendUnauthorizedResponse(c, &errMsg)
		}

		var user models.UserModel
		result := appMiddleware.DB.First(&user, claims.ID)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			errMsg := "invalid access token"
			return common.SendUnauthorizedResponse(c, &errMsg)
		}

		if result.Error != nil {
			errMsg := "invalid access token"
			return common.SendUnauthorizedResponse(c, &errMsg)
		}

		c.Set("user", user)
		return next(c)
	}
}
