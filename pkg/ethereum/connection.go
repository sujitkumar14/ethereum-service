package ethereum

import (
	"eth-service/pkg/econfig"
	"eth-service/pkg/log"

	"github.com/ethereum/go-ethereum/ethclient"
)

const BURN_ADDRESS = "0x0000000000000000000000000000000000000000"

func (e *Ethereum) CreateConnection() {

	connString := econfig.GetString("ethereum.network")
	conn, err := ethclient.Dial(connString)

	if err != nil {
		log.Debug(err)
	}

	e.conn = conn
	setInstance(e)
}

//Ping to test the Ethereum blockchain connection
//return true if connected
func (e *Ethereum) Ping() bool {

	_, err := e.GetBalance(BURN_ADDRESS)
	if err != nil {
		log.Debug(err)
		return false
	}
	return true
}
