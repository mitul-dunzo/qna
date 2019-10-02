package config

import (
	newrelic "github.com/newrelic/go-agent"
	"github.com/newrelic/go-agent/_integrations/nrlogrus"
	"github.com/sirupsen/logrus"
)

var NewRelicApp newrelic.Application

func SetupNewRelic() {
	appName := "qna"
	license := "eu01xx84b0eb9c65dcdf3b45c53380770e434f3f"
	config := newrelic.NewConfig(appName, license)
	logrus.SetLevel(logrus.DebugLevel)
	config.Logger = nrlogrus.StandardLogger()
	config.Enabled = false
	app, err := newrelic.NewApplication(config)
	if err != nil {
		logrus.Fatal("Failed to connect to NewRelic with error: ", err.Error())
		panic(err)
	}
	NewRelicApp = app
}
