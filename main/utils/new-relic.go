package utils

import (
	"github.com/gorilla/mux"
	newrelic "github.com/newrelic/go-agent"
	"github.com/newrelic/go-agent/_integrations/nrlogrus"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

var NewRelicApp newrelic.Application

func SetupNewRelic() {
	appName := os.Getenv("NewRelicAppName")
	license := os.Getenv("NewRelicLicenseKey")
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

func Instrument(v0mux *mux.Router, apiPath string, usersHandler func(w http.ResponseWriter, req *http.Request)) *mux.Route {

	return v0mux.HandleFunc(newrelic.WrapHandleFunc(NewRelicApp, apiPath, usersHandler))
}
