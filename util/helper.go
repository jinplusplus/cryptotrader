package util

import (
	"github.com/forchain/cryptotrader/model"
)

func GetNonZeroBalance(balances []model.Balance) []model.Balance {
	res := []model.Balance{}
	for _, v := range balances {
		if v.Free != 0 || v.Frozen != 0 {
			res = append(res, v)
		}
	}

	return res
}
