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

func ERC1155BalanceOf(ctx context.Context, address string, params map[string]interface{}, client *ethclient.Client, blockNumber *big.Int) *big.Float {

	addressValue, ok := params["address"]
	if !ok {
		return nil
	}
	tokenAddress, ok := addressValue.(string)
	if !ok {
		return nil
	}

	tokenIdValue, ok := params["tokenId"]
	if !ok {
		return nil
	}
	tokenIdHex, ok := tokenIdValue.(string)
	if !ok {
		return nil
	}

	// Remove the "0x" prefix if it's present
	if len(tokenIdHex) > 2 && tokenIdHex[:2] == "0x" {
		tokenIdHex = tokenIdHex[2:]
	}

	// Convert the hexadecimal string to a big.Int
	tokenId := new(big.Int)
	tokenId.SetString(tokenIdHex, 16)

	decimalsValue, ok := params["decimals"]
	if !ok {
		return nil
	}
	decimals, ok := decimalsValue.(float64)
	if !ok {
		return nil
	}

	token := token.NewToken(tokenAddress, math.Pow10(int(decimals)))
	abiName := abiUtils.TOKEN_ID
	balanceChan, errChan := token.Balance(
		ctx,
		client,
		abiUtils.GetFuncName(abiName),
		abiUtils.GetABI(abiName),
		[]interface{}{
			common.HexToAddress(address),
			tokenId,
		},
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
