package ethereumwrapper

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
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
	EthereumInstance := getInstance()
	balance, err = EthereumInstance.conn.BalanceAt(ctx, hexToAddress(address), nil)
	return
}

func (e *Ethereum) GetSyncProcess() (*ethereum.SyncProgress, error) {

	ctx := context.Background()
	EthereumInstance := getInstance()
	syncProgress, err := EthereumInstance.conn.SyncProgress(ctx)
	if syncProgress == nil {
		return nil, nil
	}
	return syncProgress, err
}

func (e *Ethereum) GetCurrentBlockNumber() (block string, err error) {

	ctx := context.Background()
	EthereumInstance := getInstance()
	header, err := EthereumInstance.conn.HeaderByNumber(ctx, nil)
	block = header.Number.String()
	return block, err
}
