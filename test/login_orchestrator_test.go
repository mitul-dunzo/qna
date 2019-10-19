package test

import (
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/suite"
	"net/http"
	"qna/main/constants"
	"qna/main/orchestrator"
	"qna/test/mocks"
	"qna/test/utils"
	"testing"
)

type LoginOrchestratorTestSuite struct {
	suite.Suite
	ctrl         *gomock.Controller
	router       *mux.Router
	userService  *mocks.MockIUserService
	otpService   *mocks.MockIOtpService
	orchestrator orchestrator.LoginOrchestrator
}

func (suite *LoginOrchestratorTestSuite) SetupTest() {
}

func (suite *LoginOrchestratorTestSuite) BeforeTest(suiteName, testName string) {
	suite.ctrl = gomock.NewController(suite.T())
	suite.userService = mocks.NewMockIUserService(suite.ctrl)
	suite.otpService = mocks.NewMockIOtpService(suite.ctrl)
	suite.orchestrator = orchestrator.NewLoginOrchestrator(suite.otpService, suite.userService)

	suite.router = mux.NewRouter().PathPrefix("/auth").Subrouter()
	suite.orchestrator.Handle(suite.router)
}

func (suite *LoginOrchestratorTestSuite) AfterTest(suiteName, testName string) {
	suite.ctrl.Finish()
}

func TestLoginOrchestratorTestSuite(t *testing.T) {
	suite.Run(t, new(LoginOrchestratorTestSuite))
}

func (suite *LoginOrchestratorTestSuite) TestLogin() {
	details := utils.NewMockUserDetails()

	suite.otpService.EXPECT().SendOtp(details).Return(nil).Times(1)

	loginUrl := constants.AuthPrefix + constants.LoginEp
	w := utils.SendPostRequest(suite.router, loginUrl, details)

	suite.Equal(w.Code, http.StatusOK)
}

func (suite *LoginOrchestratorTestSuite) TestVerifyOtp() {
	requestParams := utils.NewMockUserOtp()
	userDetails := utils.NewMockUserDetails()
	jwt := utils.NewMockJwt()

	suite.otpService.EXPECT().ValidateOtp(requestParams.PhoneNumber, requestParams.Otp).Return(userDetails, nil).Times(1)
	suite.userService.EXPECT().CreateUser(userDetails).Return(jwt, nil).Times(1)

	verifyOtpUrl := constants.AuthPrefix + constants.VerifyOtpEp
	w := utils.SendPostRequest(suite.router, verifyOtpUrl, requestParams)

	suite.Equal(w.Code, http.StatusOK)

	var response map[string]string
	_ = json.Unmarshal(w.Body.Bytes(), &response)

	suite.Equal(response["jwt"], jwt)
}
