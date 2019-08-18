package ethereum

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

type Ethereum struct {
	conn *ethclient.Client
}

func (e *Ethereum) CreateConnection() {
	
} 