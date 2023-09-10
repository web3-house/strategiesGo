package strategies

import (
	"context"
	"log"
	"math/big"
	"strconv"
	"strings"

	"github.com/This-Is-Prince/strategiesGo/token"
	abiUtils "github.com/This-Is-Prince/strategiesGo/utils/abi"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ERC721WithTokenId(ctx context.Context, address string, params map[string]interface{}, client *ethclient.Client, blockNumber *big.Int) *big.Float {

	addressValue, ok := params["address"]
	if !ok {
		return nil
	}
	tokenAddress, ok := addressValue.(string)
	if !ok {
		return nil
	}

	tokenIdsValue, ok := params["tokenIds"]
	if !ok {
		return nil
	}
	tokenIdInterfaces, ok := tokenIdsValue.([]interface{})
	if !ok {
		return nil
	}

	tokenIds := []*big.Int{}
	for _, tokenIdInterface := range tokenIdInterfaces {
		tokenIdValue, ok := tokenIdInterface.(string)
		if ok {
			tokenIdNum, err := strconv.ParseInt(tokenIdValue, 10, 64)
			if err != nil {
				log.Println("Error:-", err)
				continue
			}
			uintValue := big.NewInt(tokenIdNum)
			tokenIds = append(tokenIds, uintValue)
		}
	}

	token := token.NewToken(tokenAddress, 0)
	abiName := abiUtils.OWNER_OF

	for _, tokenId := range tokenIds {
		ownerChan, errChan := token.OwnerOf(
			ctx,
			client,
			abiUtils.GetFuncName(abiName),
			abiUtils.GetABI(abiName),
			[]interface{}{tokenId},
			blockNumber,
		)

		select {
		case owner := <-ownerChan:
			if strings.ToLower(owner.String()) == strings.ToLower(address) {
				return new(big.Float).SetInt64(1)
			}
		case err := <-errChan:
			log.Println(err)
			return nil
		}
	}

	return new(big.Float).SetInt64(0)
}
