package services

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"qna/main/constants"
	"qna/main/dtos"
)

type IUserService interface {
	CreateUser(details *dtos.UserDetails) (string, error)
}

type UserService struct {
	db  *gorm.DB
	jwt IJwtService
}

func NewUserService(db *gorm.DB, jwt IJwtService) IUserService {
	return &UserService{
		db:  db,
		jwt: jwt,
	}
}

func (service *UserService) CreateUser(details *dtos.UserDetails) (string, error) {
	user := &dtos.User{
		UserDetails: details,
	}
	err := service.db.Where(constants.PhoneNumberQuery, details.PhoneNumber).FirstOrCreate(user).Error
	if err != nil {
		logrus.Error("Couldn't find user: ", err.Error())
		return "", err
	}

	token, err := service.jwt.CreateToken(user.ID)
	if err != nil {
		logrus.Error("Problem in creating jwt: ", err.Error())
		return "", err
	}

	return token, nil
}
