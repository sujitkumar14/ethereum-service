package ethereum

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

type Ethereum struct {
	conn *ethclient.Client
}

var EthereumInstance *Ethereum

func New() *Ethereum {
	return &Ethereum{}
}

func getInstance() *Ethereum {
	return EthereumInstance
}

func setInstance(etheremInstance *Ethereum) {
	EthereumInstance = etheremInstance
}
