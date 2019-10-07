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

	user, err := service.getUser(details)
	if err != nil {
		return "", err
	}

	token, err := service.jwt.CreateToken(user.ID)
	if err != nil {
		logrus.Error("Problem in creating jwt", err.Error())
		service.deleteUser(user)
		return "", err
	}

	user.Jwt = token
	service.db.Save(user)
	return token, nil
}

func (service *UserService) getUser(details *dtos.UserDetails) (*dtos.User, error) {
	user := &dtos.User{}
	err := service.db.First(user, &dtos.User{
		UserDetails: &dtos.UserDetails{
			PhoneNumber: details.PhoneNumber,
		},
	}).Error
	if err == nil {
		return user, nil
	}

	user = &dtos.User{
		UserDetails: details,
	}
	err = service.db.Create(&user).Error
	if err != nil {
		logrus.Error("User creation failed: ", err.Error())
		return nil, errors.New("user creation failed")
	}

	return user, nil
}

func (service *UserService) deleteUser(user *dtos.User) {
	service.db.Delete(user)
}
