package account

import (
	"context"
	AccountProto "eth-service/api/account"
)

type Account struct{}

func (a *Account) getBalance(ctx context.Context, request *AccountProto.BalanceRequest) (response *AccountProto.BalanceResponse, err error) {

}
