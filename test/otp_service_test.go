package test

import (
	"github.com/go-redis/redis"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"qna/main/dtos"
	"qna/main/services"
	"qna/test/mocks"
	"testing"
	"time"
)

type OtpServiceTestSuite struct {
	suite.Suite
	ctrl           *gomock.Controller
	smsClient      *mocks.MockISmsClient
	randNumService *mocks.MockIRandNumService
	redis          *redis.Client
	service        services.OtpService
}

func (suite *OtpServiceTestSuite) SetupTest() {
}

func (suite *OtpServiceTestSuite) BeforeTest(suiteName, testName string) {
	suite.ctrl = gomock.NewController(suite.T())
	suite.smsClient = mocks.NewMockISmsClient(suite.ctrl)
	suite.redis = mocks.NewMockRedis()
	suite.randNumService = mocks.NewMockIRandNumService(suite.ctrl)
	suite.service = services.NewOtpService(suite.redis, suite.smsClient, suite.randNumService)
}

func (suite *OtpServiceTestSuite) AfterTest(suiteName, testName string) {
	mocks.StopMockRedis()
	suite.ctrl.Finish()
}

func TestOtpServiceTestSuite(t *testing.T) {
	suite.Run(t, new(OtpServiceTestSuite))
}

func (suite *OtpServiceTestSuite) TestSendOtp() {
	userDetails := dtos.UserDetails{
		PhoneNumber: "9876543210",
		Name:        "Test User",
		Email:       "Test@dunzo.in",
	}
	otp := "1234"

	suite.randNumService.EXPECT().GetRandNum().Return(otp, nil).Times(1)
	suite.smsClient.EXPECT().SendOtpSms(otp, userDetails.PhoneNumber).Return(nil).Times(1)

	err := suite.service.SendOtp(&userDetails)
	suite.Nil(err)
}

func (suite *OtpServiceTestSuite) TestValidateOtp() {
	otp := "1234"
	userDetails := dtos.UserDetails{
		PhoneNumber: "9876543210",
		Name:        "Test User",
		Email:       "Test@dunzo.in",
	}
	otpDetails := dtos.UserOtp{
		UserDetails: &userDetails,
		Otp:         otp,
	}

	_ = suite.redis.Set(userDetails.PhoneNumber, otpDetails, 3*time.Minute)
	details, err := suite.service.ValidateOtp(userDetails.PhoneNumber, otp)

	suite.Nil(err)
	suite.Equal(details, &userDetails)
}
