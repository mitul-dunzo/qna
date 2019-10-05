package services

import (
	"crypto/rand"
	"github.com/sirupsen/logrus"
	"math/big"
	"qna/main/clients"
	"qna/main/config"
	"strconv"
	"time"
)

func SendOtp(phoneNumber string) error {
	otp, err := getRandNum()
	if err != nil {
		logrus.Error("Couldn't create OTP: ", err.Error())
		return err
	}

	err = config.Save(phoneNumber, otp, 3*time.Minute)
	if err != nil {
		logrus.Error("Couldn't save number in redis")
		return err
	}

	err = clients.SendOtpSms(otp, phoneNumber)
	if err != nil {
		logrus.Error("Couldn't send OTP: ", err.Error())
		return err
	}

	return nil
}

func getRandNum() (string, error) {
	nBig, e := rand.Int(rand.Reader, big.NewInt(9999))
	if e != nil {
		logrus.Error("Couldn't generate a random number: ", e.Error())
		return "", e
	}
	return strconv.FormatInt(nBig.Int64()+1000, 10), nil
}
