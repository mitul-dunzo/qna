package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
)

var dbConn *gorm.DB

func SetupDatabase() {
	user := "qna_admin"
	password := "admin_1234"
	database := "qna"
	host := "0.0.0.0"
	port := 5432
	psqlConf := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, database)

	logrus.Info("Connecting to : {}", psqlConf)

	db, err := gorm.Open("postgres", psqlConf)
	if err != nil {
		logrus.Fatal("failed to connect database with error: ", err.Error())
	}
	maxIdleConn := 10
	maxOpenConn := 10
	logMode := true
	db.LogMode(logMode)
	dbConn = db
	db.DB().SetMaxIdleConns(maxIdleConn)
	db.DB().SetMaxOpenConns(maxOpenConn)
}
