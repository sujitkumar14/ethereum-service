package account

import (
	"context"
	AccountProto "eth-service/api/account"
	"eth-service/pkg/helper"
)

type Account struct{}

func (a *Account) GetBalance(ctx context.Context, request *AccountProto.BalanceRequest) (response *AccountProto.BalanceResponse, err error) {

	var balance string
	var AccountHelper *helper.Account
	var address string

	address = request.GetAddress()
	AccountHelper = helper.New()
	balance, err = AccountHelper.GetBalance(ctx, address)
	if err != nil {
		return &AccountProto.BalanceResponse{
			Response: &AccountProto.Response{
				Status: AccountProto.Status_FAILURE,
			},
		}, nil
	} else {

		return &AccountProto.BalanceResponse{
			Response: &AccountProto.Response{
				Status: AccountProto.Status_SUCCESS,
			},
			Balance: balance,
		}, nil
	}

}
