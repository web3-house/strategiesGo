package strategies

import (
	"context"
	"log"
	"math"
	"math/big"

	"github.com/This-Is-Prince/strategiesGo/token"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ERC20BalanceOf(ctx context.Context, address string, params map[string]interface{}, client *ethclient.Client, blockNumber *big.Int) *big.Float {

	addressValue, ok := params["address"]
	if !ok {
		return nil
	}
	tokenAddress, ok := addressValue.(string)
	if !ok {
		return nil
	}
	decimalsValue, ok := params["decimals"]
	if !ok {
		return nil
	}
	decimals, ok := decimalsValue.(float64)
	if !ok {
		return nil
	}
	token := token.NewToken(tokenAddress, math.Pow10(int(decimals)))
	balanceChan, errChan := token.Balance(ctx, client, common.HexToAddress(address), blockNumber)

	select {
	case balance := <-balanceChan:
		return token.Format(balance)
	case err := <-errChan:
		log.Println(err)
	}

	return nil
}
