package main

import (
	Econfig "eth-service/pkg/econfig"
	"eth-service/pkg/ethereum"
	"eth-service/pkg/log"
)

func main() {

	initializeConfig()
	connectWithBlockchain()
}

func connectWithBlockchain() {

	Ethereum := ethereum.New()
	Ethereum.CreateConnection()
	if !Ethereum.Ping() {
		log.Debug("Failed to establish ethereum connection")
	}
	log.Debug("ethereum Connected")
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
