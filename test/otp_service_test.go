package test

import (
	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"qna/main/services"
	"qna/test/mocks"
	"qna/test/utils"
	"testing"
	"time"
)

type OtpServiceTestSuite struct {
	suite.Suite
	ctrl           *gomock.Controller
	smsClient      *mocks.MockISmsClient
	randNumService *mocks.MockIRandNumService
	redis          *redis.Client
	miniRedis      *miniredis.Miniredis
	service        services.IOtpService
}

func (suite *OtpServiceTestSuite) SetupTest() {
}

func (suite *OtpServiceTestSuite) BeforeTest(suiteName, testName string) {
	suite.ctrl = gomock.NewController(suite.T())
	suite.smsClient = mocks.NewMockISmsClient(suite.ctrl)
	suite.redis, suite.miniRedis = mocks.NewMockRedis()
	suite.randNumService = mocks.NewMockIRandNumService(suite.ctrl)
	suite.service = services.NewOtpService(suite.redis, suite.smsClient, suite.randNumService)
}

func (suite *OtpServiceTestSuite) AfterTest(suiteName, testName string) {
	suite.miniRedis.Close()
	suite.ctrl.Finish()
}

func TestOtpServiceTestSuite(t *testing.T) {
	suite.Run(t, new(OtpServiceTestSuite))
}

func (suite *OtpServiceTestSuite) TestSendOtp() {
	userDetails := utils.NewMockUserDetails()
	otp := utils.NewMockOtp()

	suite.randNumService.EXPECT().GetRandNum().Return(otp, nil).Times(1)
	suite.smsClient.EXPECT().SendOtpSms(otp, userDetails.PhoneNumber).Return(nil).Times(1)

	err := suite.service.SendOtp(userDetails)
	suite.Nil(err)
}

func (suite *OtpServiceTestSuite) TestValidateOtp() {
	otp := utils.NewMockOtp()
	userDetails := utils.NewMockUserDetails()
	userOtp := utils.NewMockUserOtp()

	_ = suite.redis.Set(userDetails.PhoneNumber, userOtp, 3*time.Minute)
	details, err := suite.service.ValidateOtp(userDetails.PhoneNumber, otp)

	suite.Nil(err)
	suite.Equal(details, userDetails)
}
