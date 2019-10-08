package dtos

import "github.com/jinzhu/gorm"

type UserDetails struct {
	PhoneNumber string `json:"phone_number"`
	Name        string `json:"name"`
	Email       string `json:"email"`
}

type User struct {
	gorm.Model
	*UserDetails
	ID  uint
	Jwt string
}

func (User) TableName() string {
	return "users"
}
