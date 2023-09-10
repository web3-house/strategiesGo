package strategies

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func ERC721Enumerable(ctx context.Context, address string, params map[string]interface{}, client *ethclient.Client, blockNumber *big.Int) *big.Float {
	params["decimals"] = float64(0)
	return ERC20BalanceOf(ctx, address, params, client, blockNumber)
}
