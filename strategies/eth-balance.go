package strategies

import (
	"context"
	"log"
	"math"
	"math/big"

	"github.com/This-Is-Prince/strategiesGo/token"
	abiUtils "github.com/This-Is-Prince/strategiesGo/utils/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func EthBalance(ctx context.Context, address string, params map[string]interface{}, client *ethclient.Client, blockNumber *big.Int) *big.Float {
	addressValue, ok := params["address"]
	if !ok {
		return nil
	}
	tokenAddress, ok := addressValue.(string)
	if !ok {
		return nil
	}

	token := token.NewToken(tokenAddress, math.Pow10(18))
	abiName := abiUtils.GET_ETH_BALANCE
	balanceChan, errChan := token.Balance(
		ctx,
		client,
		abiUtils.GetFuncName(abiName),
		abiUtils.GetABI(abiName),
		[]interface{}{common.HexToAddress(address)},
		blockNumber,
	)

	select {
	case balance := <-balanceChan:
		return token.Format(balance)
	case err := <-errChan:
		log.Println(err)
	}

	return nil
}
