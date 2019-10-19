package dtos

type OtpValidationSuccess struct {
	Jwt string `json:"jwt"`
}

type UserOtp struct {
	*UserDetails
	Otp string `json:"otp"`
}
