package services

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"os"
	"qna/main/constants"
	"qna/main/dtos"
	"time"
)

var jwtKey = []byte(os.Getenv("JWTKey"))

type IJwtService interface {
	CreateToken(id uint) (string, error)
	ValidateUser(tokenString string) (uint, error)
}

type JwtService struct {
}

func NewJwtService() IJwtService {
	return &JwtService{}
}

func (service *JwtService) CreateToken(id uint) (string, error) {
	claims := &dtos.Claims{
		UserId: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: (time.Now().Add(2 * 24 * time.Hour)).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		logrus.Error("Couldn't create tokenString: ", err.Error())
		return "", err
	}

	return tokenString, nil
}

func (service *JwtService) ValidateUser(tokenString string) (uint, error) {
	token, err := jwt.ParseWithClaims(tokenString, &dtos.Claims{}, func(token *jwt.Token) (i interface{}, e error) {
		return jwtKey, nil
	})
	if err != nil {
		logrus.Debug("Error in parsing claims")
		return 0, err
	}

	claims, ok := token.Claims.(*dtos.Claims)
	if !ok || !token.Valid {
		return 0, constants.WrongTokenError
	}

	return claims.UserId, nil
}
