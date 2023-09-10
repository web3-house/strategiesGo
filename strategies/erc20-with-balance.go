package strategies

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func ERC20WithBalance(ctx context.Context, address string, params map[string]interface{}, client *ethclient.Client, blockNumber *big.Int) *big.Float {
	minBalanceValue, ok := params["minBalance"]
	if !ok {
		minBalanceValue = float64(0)
	}
	minBalance, ok := minBalanceValue.(float64)
	if !ok {
		minBalance = float64(0)
	}

	balance := ERC20BalanceOf(ctx, address, params, client, blockNumber)
	if balance != nil {
		if balance.Cmp(new(big.Float).SetFloat64(minBalance)) >= 1 {
			return big.NewFloat(1)
		} else {
			return big.NewFloat(0)
		}
	}
	return nil
}
