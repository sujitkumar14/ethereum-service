package helper

import (
	"context"
	ethereum "eth-service/pkg/ethereum"
	"eth-service/pkg/log"
	"math/big"
)

type Eth struct{}

func New() *Eth {
	return &Eth{}
}

func (e *Eth) GetBalance(ctx context.Context, address string) (balance string, err error) {

	var Ethereum *ethereum.Ethereum
	var balanceInt *big.Int
	Ethereum = ethereum.New()

	balanceInt, err = Ethereum.GetBalance(address)
	if err != nil {
		log.Debug(err)
	}
	balance = balanceInt.String()
	return

}

func (e *Eth) GetBlockNumber(ctx context.Context) (blockNumber string, err error) {
	var Ethereum *ethereum.Ethereum

	Ethereum = ethereum.New()

	blockNumber, err = Ethereum.GetCurrentBlockNumber()

	if err != nil {
		log.Debug(err)
	}
	return blockNumber, err
}
