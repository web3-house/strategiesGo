package strategies

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func ERC721WithMultiplier(ctx context.Context, address string, params map[string]interface{}, client *ethclient.Client, blockNumber *big.Int) *big.Float {

	multiplierValue, ok := params["multiplier"]
	if !ok {
		multiplierValue = float64(0)
	}
	multiplier, ok := multiplierValue.(float64)
	if !ok {
		multiplier = float64(0)
	}

	balance := ERC721(ctx, address, params, client, blockNumber)
	if balance != nil {
		return balance.Mul(balance, new(big.Float).SetFloat64(multiplier))
	}

	return nil
}
