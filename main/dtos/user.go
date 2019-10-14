package dtos

type UserDetails struct {
	PhoneNumber string `json:"phone_number"`
	Name        string `json:"name"`
	Email       string `json:"email"`
}

type User struct {
	*UserDetails
	ID uint `gorm:"PRIMARY_KEY"`
}

func (User) TableName() string {
	return "users"
}
