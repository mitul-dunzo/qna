package test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/suite"
	"qna/main/services"
	"qna/test/mocks"
	"qna/test/utils"
	"regexp"
	"testing"
)

type UserServiceTestSuite struct {
	suite.Suite
	ctrl       *gomock.Controller
	db         *gorm.DB
	dbMocker   sqlmock.Sqlmock
	jwtService *mocks.MockIJwtService
	service    services.IUserService
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
	userDetails := utils.NewMockUserDetails()
	id := utils.NewMockUserID()

	query := utils.NewUserQuery()
	newRow := utils.NewUserTableRow()
	newRow.AddRow(id, userDetails.Name, userDetails.PhoneNumber, userDetails.Email)
	suite.dbMocker.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(userDetails.PhoneNumber).WillReturnRows(newRow)

	token := utils.NewMockJwt()
	suite.jwtService.EXPECT().CreateToken(id).Return(token, nil).Times(1)

	returnToken, err := suite.service.CreateUser(userDetails)
	suite.Nil(err)
	suite.Equal(returnToken, token)
}
