package main

import (
	"context"
	server "eth-service/api/account"
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

	client := server.NewAccountClient(conn)
	response, err := client.GetBalance(context.Background(), &server.BalanceRequest{
		Address: "0x0000000000000000000000000000000000000000",
	})
	fmt.Println(response)
}
