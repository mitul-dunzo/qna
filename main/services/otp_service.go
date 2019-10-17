package services

import (
	"errors"
	r "github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"qna/main/clients"
	"qna/main/dtos"
	"time"
)

type OtpService struct {
	redis          *r.Client
	smsClient      clients.ISmsClient
	randNumService IRandNumService
}

const InvalidOtp = "invalid otp"

func NewOtpService(redis *r.Client, smsClient clients.ISmsClient, randNumService IRandNumService) OtpService {
	return OtpService{
		redis:          redis,
		smsClient:      smsClient,
		randNumService: randNumService,
	}
}

func (service *OtpService) SendOtp(details *dtos.UserDetails) error {
	otp, err := service.randNumService.GetRandNum()
	if err != nil {
		logrus.Error("Couldn't create OTP: ", err.Error())
		return err
	}

	otpDetails := dtos.UserOtp{
		UserDetails: details,
		Otp:         otp,
	}

	err = service.redis.Set(details.PhoneNumber, otpDetails, 3*time.Minute).Err()
	if err != nil {
		logrus.Error("Couldn't save otp in redis: ", err.Error())
		return err
	}

	err = service.smsClient.SendOtpSms(otp, details.PhoneNumber)
	if err != nil {
		logrus.Error("Couldn't send sms: ", err.Error())
		return err
	}

	logrus.Debug("OTP is ", otp)

	return nil
}

func (service *OtpService) ValidateOtp(phoneNumber string, otp string) (*dtos.UserDetails, error) {
	var details dtos.UserOtp
	err := service.redis.Get(phoneNumber).Scan(&details)
	if err != nil {
		logrus.Error("No details of the user in redis: ", err.Error())
		return nil, err
	}

	if otp != details.Otp {
		logrus.Error("OTP didn't match")
		return nil, errors.New(InvalidOtp)
	}

	return details.UserDetails, nil
}
