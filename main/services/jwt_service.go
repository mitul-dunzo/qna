package services

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"qna/main/dtos"
	"time"
)

var jwtKey = []byte("7730DE6E5C013DC1C64E3DBE791460CE88C06D4B970EADD70480EE46E4CFE60B")

type JwtService struct {
}

func NewJwtService() JwtService {
	return JwtService{}
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

	logrus.Debug("JWT is ", tokenString)

	return tokenString, nil
}

func (service *JwtService) ValidateUser(tokenString string) (uint, error) {
	token, err := jwt.ParseWithClaims(tokenString, &dtos.Claims{}, func(token *jwt.Token) (i interface{}, e error) {
		return jwtKey, nil
	})
	if err != nil {
		logrus.Debug("Error is parsing claims")
		return 0, err
	}

	claims, ok := token.Claims.(*dtos.Claims)
	if !ok || !token.Valid {
		logrus.Error("toke not valid")
		return 0, errors.New("wrong token")
	}

	logrus.Debug("User id is: ", claims.UserId)

	return claims.UserId, nil
}
