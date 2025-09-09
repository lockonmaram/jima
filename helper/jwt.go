package helper

import (
	"jima/config"
	"jima/entity/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Serial   string `json:"serial"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJWT(config config.Config, user *model.User) (signedToken string, err error) {
	claims := &Claims{
		Serial:   user.Serial,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * time.Duration(config.JWTExpireDays))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err = unsignedToken.SignedString([]byte(config.JWTSecret))

	return signedToken, err
}

func ValidateJWT(config config.Config, signedToken string) (claims *Claims, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(config.JWTSecret), nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, err
	}

	if claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, jwt.ErrTokenExpired
	}

	return claims, nil
}
