package test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/suite"
	"qna/main/dtos"
	"qna/main/services"
	"qna/test/mocks"
	"regexp"
	"testing"
)

type UserServiceTestSuite struct {
	suite.Suite
	ctrl       *gomock.Controller
	db         *gorm.DB
	dbMocker   sqlmock.Sqlmock
	jwtService *mocks.MockIJwtService
	service    services.UserService
}

func (suite *UserServiceTestSuite) SetupTest() {
}

func (suite *UserServiceTestSuite) BeforeTest(suiteName, testName string) {
	suite.ctrl = gomock.NewController(suite.T())
	suite.db, suite.dbMocker = mocks.GetMockDB()
	suite.jwtService = mocks.NewMockIJwtService(suite.ctrl)
	suite.service = services.NewUserService(suite.db, suite.jwtService)
}

func (suite *UserServiceTestSuite) AfterTest(suiteName, testName string) {
	if err := suite.dbMocker.ExpectationsWereMet(); err != nil {
		suite.T().Errorf("there were unfulfilled expectations: %s", err)
	}
	suite.ctrl.Finish()
}

func TestUserServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}

func (suite *UserServiceTestSuite) TestCreateUser() {
	userDetails := dtos.UserDetails{
		PhoneNumber: "9876543210",
		Name:        "Test User",
		Email:       "test@dunzo.in",
	}
	id := uint(123)

	query := "SELECT * FROM \"users\"  WHERE (phone_number = $1) ORDER BY \"users\".\"id\" ASC LIMIT 1"
	newRow := sqlmock.NewRows([]string{"id", "name", "phone_number", "email"})
	newRow.AddRow(id, userDetails.Name, userDetails.PhoneNumber, userDetails.Email)
	suite.dbMocker.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(userDetails.PhoneNumber).WillReturnRows(newRow)

	token := "Random JWT token String"
	suite.jwtService.EXPECT().CreateToken(id).Return(token, nil).Times(1)

	returnToken, err := suite.service.CreateUser(&userDetails)
	suite.Nil(err)
	suite.Equal(returnToken, token)
}
