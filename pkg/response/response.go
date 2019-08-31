package response

import (
	EthProto "eth-service/api/eth"
)

type Response struct{}

func New() *Response {
	return &Response{}
}

func (r *Response) CurrentBlockNumberResponse(status EthProto.Status, blocknumber string) (respnse *EthProto.CurrentBlockNumberResponse, err error) {

	return &EthProto.CurrentBlockNumberResponse{
		Response: &EthProto.Response{
			Status: status,
		},
		Block: blocknumber,
	}, nil
}

func (r *Response) GetBalanceResponse(status EthProto.Status, balance string) (response *EthProto.BalanceResponse, err error) {

	return &EthProto.BalanceResponse{
		Response: &EthProto.Response{
			Status: status,
		},
		Balance: balance,
	}, nil

}
