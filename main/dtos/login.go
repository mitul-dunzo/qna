package dtos

import "encoding/json"

type OtpValidationSuccess struct {
	Jwt string `json:"jwt"`
}

type UserOtp struct {
	*UserDetails
	Otp string `json:"otp"`
}

func (u UserOtp) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}

func (u *UserOtp) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)
}
