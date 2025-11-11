package common

import (
	"errors"
	"os"
	"time"

	"github.com/Dunsin-cyber/bkeeper/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/gommon/log"
)

type CustomJWTClaims struct {
	ID uint `json:"id"`
	jwt.RegisteredClaims
}

func GenerateJWT(user models.UserModel) (*string, *string, error) {
	userClaims := CustomJWTClaims{
		ID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // Token expires in 24 hours
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	signedAccessToken, err := accessToken.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	if err != nil {
		return nil, nil, err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &CustomJWTClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // Token expires in 24 hours
		},
	})
	signedRefreshToken, err := refreshToken.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	if err != nil {
		return nil, nil, err
	}

	return &signedAccessToken, &signedRefreshToken, nil
}

func ParseJWTSignedAccessToken(tokenString string) (*CustomJWTClaims, error) {
	parseJwtAccessToken, err := jwt.ParseWithClaims(tokenString, &CustomJWTClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		log.Error(err)
		return nil, err
	} else if claims, ok := parseJwtAccessToken.Claims.(*CustomJWTClaims); ok && parseJwtAccessToken.Valid {
		return claims, nil
	} else {
		return nil, errors.New("unknown claims type, cannot proceed")
	}

}

func IsClaimExpired(claims *CustomJWTClaims) bool {
	currrentTime := jwt.NewNumericDate(time.Now())
	return claims.ExpiresAt.Before(currrentTime.Time)
}
