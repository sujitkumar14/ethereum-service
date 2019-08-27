package helper

import (
	"context"
	"eth-service/pkg/ethereum"
	"eth-service/pkg/log"
	"math/big"
)

type Account struct{}

func New() *Account {
	return &Account{}
}

func (a *Account) GetBalance(ctx context.Context, address string) (balance string, err error) {

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
