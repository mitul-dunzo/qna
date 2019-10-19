package utils

import "qna/main/dtos"

func NewMockUserID() uint {
	return 123
}

func NewMockOtp() string {
	return "1234"
}

func NewMockUserDetails() *dtos.UserDetails {
	return &dtos.UserDetails{
		PhoneNumber: "9876543210",
		Name:        "Test User",
		Email:       "Test@dunzo.in",
	}
}

func NewMockUserOtp() *dtos.UserOtp {
	return &dtos.UserOtp{
		UserDetails: NewMockUserDetails(),
		Otp:         NewMockOtp(),
	}
}

func NewMockJwt() string {
	return "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzE2NDIwOTQsInVzZXJfaWQiOjV9.laBxN2sm74y_E5EGjBfqM6RSwg1fmZSrJvW6-1vpOg0"
}
