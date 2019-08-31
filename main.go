package main

import (
	EthProto "eth-service/api/eth"
	EthController "eth-service/internal/controller/eth"
	Econfig "eth-service/pkg/econfig"
	ethereum "eth-service/pkg/ethereum"
	"eth-service/pkg/log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	initializeConfig()
	connectWithBlockchain()
	startGrpcServer()
}

func startGrpcServer() {

	lis, err := net.Listen("tcp", "localhost:4000")
	if err != nil {
		log.Debug(err)
	}
	grpcServer := grpc.NewServer()
	EthProto.RegisterAccountServer(grpcServer, &EthController.Eth{})
	grpcServer.Serve(lis)
	log.Debug("Its working")
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
