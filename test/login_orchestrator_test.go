package test

import (
	"bytes"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"qna/main/dtos"
	"qna/main/orchestrator"
	"qna/test/mocks"
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
	details := dtos.UserDetails{
		PhoneNumber: "9876543210",
		Name:        "Test User",
		Email:       "test@dunzo.in",
	}

	suite.otpService.EXPECT().SendOtp(&details).Return(nil).Times(1)

	data, _ := json.Marshal(details)
	r, _ := http.NewRequest(http.MethodPost, "/auth/login", bytes.NewBuffer(data))

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, r)

	suite.Equal(w.Code, http.StatusOK)
}

func (suite *LoginOrchestratorTestSuite) TestVerifyOtp() {
	requestParams := dtos.OtpData{
		Otp:         "1234",
		PhoneNumber: "9876543210",
	}

	userDetails := dtos.UserDetails{
		PhoneNumber: requestParams.PhoneNumber,
		Name:        "Test User",
		Email:       "test@dunzo.in",
	}

	jwt := "Random jwt string"

	suite.otpService.EXPECT().ValidateOtp(requestParams.PhoneNumber, requestParams.Otp).Return(&userDetails, nil).Times(1)
	suite.userService.EXPECT().CreateUser(&userDetails).Return(jwt, nil).Times(1)

	data, _ := json.Marshal(requestParams)
	r, _ := http.NewRequest(http.MethodPost, "/auth/verify-otp", bytes.NewBuffer(data))

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, r)

	suite.Equal(w.Code, http.StatusOK)

	var response map[string]string
	_ = json.Unmarshal(w.Body.Bytes(), &response)

	suite.Equal(response["jwt"], jwt)
}
