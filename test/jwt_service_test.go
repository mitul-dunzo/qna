package test

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"qna/main/services"
	"testing"
)

type JwtServiceTestSuite struct {
	suite.Suite
	ctrl    *gomock.Controller
	service services.JwtService
}

func (suite *JwtServiceTestSuite) SetupTest() {
}

func (suite *JwtServiceTestSuite) BeforeTest(suiteName, testName string) {
	suite.ctrl = gomock.NewController(suite.T())
}

func (suite *JwtServiceTestSuite) AfterTest(suiteName, testName string) {
	suite.ctrl.Finish()
}

func TestJwtServiceTestSuite(t *testing.T) {
	suite.Run(t, new(JwtServiceTestSuite))
}

func (suite *JwtServiceTestSuite) TestService() {
	var userId uint = 1234
	token, err := suite.service.CreateToken(userId)
	suite.Nil(err)

	returnedId, err := suite.service.ValidateUser(token)
	suite.Nil(err)

	suite.Equal(userId, returnedId)
}
