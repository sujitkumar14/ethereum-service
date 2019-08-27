package main

import (
	proto "eth-service/api/account"
	controller "eth-service/internal/controller/account"
	Econfig "eth-service/pkg/econfig"
	"eth-service/pkg/ethereum"
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
	proto.RegisterAccountServer(grpcServer, &controller.Account{})
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
