package services

import (
	"crypto/rand"
	"github.com/sirupsen/logrus"
	"math/big"
	"strconv"
)

type IRandNumService interface {
	GetRandNum() (string, error)
}

type RandNumService struct{}

func NewRandNumService() IRandNumService {
	return &RandNumService{}
}

func (s *RandNumService) GetRandNum() (string, error) {
	nBig, e := rand.Int(rand.Reader, big.NewInt(8999))
	if e != nil {
		logrus.Error("Couldn't generate a random number: ", e.Error())
		return "", e
	}
	return strconv.FormatInt(nBig.Int64()+1000, 10), nil
}
