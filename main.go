package main

import (
	Econfig "eth-service/pkg/econfig"
	"eth-service/pkg/log"
)

func main() {

	initializeConfig()
}

func initializeConfig() {

	Econfig.Init()
	Econfig.SetConfigName("config")
	Econfig.AddConfigPath("./configs/")
	err := Econfig.ReadInConfig()
	if err != nil {
		log.Debug(err)
	}

}
