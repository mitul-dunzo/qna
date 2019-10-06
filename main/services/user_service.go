package services

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"qna/main/dtos"
)

type UserService struct {
	db  *gorm.DB
	jwt *JwtService
}

func NewUserService(db *gorm.DB, jwt *JwtService) UserService {
	return UserService{
		db:  db,
		jwt: jwt,
	}
}

func (service *UserService) CreateUser(details *dtos.UserDetails) (string, error) {
	user := dtos.User{
		UserDetails: details,
		ID:          0,
	}

	service.db.Create(&user)
	if user.ID == 0 {
		return "", errors.New("user creation failed")
	}

	token, err := service.jwt.CreateToken(user.ID)
	if err != nil {
		logrus.Error("Problem in creating jwt", err.Error())
		service.deleteUser(&user)
		return "", err
	}

	user.Jwt = token
	service.db.Save(&user)
	return token, nil
}

func (service *UserService) deleteUser(user *dtos.User) {
	service.db.Delete(user)
}
