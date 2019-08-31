package main

import (
	"context"
	EthServer "eth-service/api/eth"
	"fmt"

	"google.golang.org/grpc"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial("localhost:4000", opts...)

	if err != nil {
		fmt.Printf("failure while dialing: %v", err)
	}
	defer conn.Close()

	// client := server.NewAccountClient(conn)
	ethClient := EthServer.NewAccountClient(conn)
	// response, err := ethClient.GetBalance(context.Background(), &EthServer.BalanceRequest{
	// 	Address: "0x0000000000000000000000000000000000000000",
	// })
	response, err := ethClient.CurrentBlockNumber(context.Background(), &EthServer.CurrentBlockNumberRequest{})
	fmt.Println(response)
}
