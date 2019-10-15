package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
	"os"
)

var db *gorm.DB

func SetupDatabase() {
	user := os.Getenv("PostgresUser")
	password := os.Getenv("PostgresPassword")
	database := os.Getenv("PostgresDBName")
	host := os.Getenv("PostgresHost")
	port := os.Getenv("PostgresPort")
	psqlConf := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, database)

	logrus.Info("Connecting to : {}", psqlConf)

	dbConn, err := gorm.Open("postgres", psqlConf)
	if err != nil {
		logrus.Fatal("failed to connect database with error: ", err.Error())
	}
	db = dbConn
}

func GetDB() *gorm.DB {
	return db
}
