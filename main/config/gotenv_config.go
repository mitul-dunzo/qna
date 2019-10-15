package config

import "github.com/subosito/gotenv"

func SetupEnv() {
	_ = gotenv.Load()
}
