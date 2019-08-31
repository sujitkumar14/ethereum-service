package eth

import (
	"context"
	EthProto "eth-service/api/eth"
	"eth-service/pkg/helper"
	responseHelper "eth-service/pkg/response"
)

type Eth struct{}

func (e *Eth) CurrentBlockNumber(context context.Context, request *EthProto.CurrentBlockNumberRequest) (response *EthProto.CurrentBlockNumberResponse, err error) {
	var blockNumber string
	var EthHelper *helper.Eth = helper.New()
	var EthResponse *responseHelper.Response = responseHelper.New()
	blockNumber, err = EthHelper.GetBlockNumber(context)

	if err == nil {
		return EthResponse.CurrentBlockNumberResponse(EthProto.Status_SUCCESS, blockNumber)
	} else {
		return EthResponse.CurrentBlockNumberResponse(EthProto.Status_FAILURE, "0")
	}
}

func (e *Eth) GetBalance(ctx context.Context, request *EthProto.BalanceRequest) (response *EthProto.BalanceResponse, err error) {

	var balance string
	var EthHelper *helper.Eth
	var address string
	var EthResponse *responseHelper.Response = responseHelper.New()

	address = request.GetAddress()
	EthHelper = helper.New()
	balance, err = EthHelper.GetBalance(ctx, address)
	if err == nil {
		return EthResponse.GetBalanceResponse(EthProto.Status_FAILURE, balance)
	} else {
		return EthResponse.GetBalanceResponse(EthProto.Status_SUCCESS, balance)
	}

}
