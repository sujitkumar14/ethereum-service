package ethereum

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func AddressToHex(address common.Address) string {
	return address.String()
}

func hexToAddress(address string) common.Address {
	return common.HexToAddress(address)
}

func (e *Ethereum) GetBalance(address string) (balance *big.Int, err error) {

	ctx := context.Background()
	balance, err = e.conn.BalanceAt(ctx, hexToAddress(address), nil)
	return
}
