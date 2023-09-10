package strategies

import (
	"context"
	"encoding/json"
	"log"
	"math"
	"math/big"

	"github.com/This-Is-Prince/strategiesGo/token"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ContractCall(ctx context.Context, address string, params map[string]interface{}, client *ethclient.Client, blockNumber *big.Int) *big.Float {
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
	methodABIValue, ok := params["methodABI"]
	if !ok {
		return nil
	}
	methodABI, ok := methodABIValue.(map[string]interface{})
	if !ok {
		return nil
	}
	nameValue, ok := methodABI["name"]
	if !ok {
		return nil
	}
	name, ok := nameValue.(string)
	if !ok {
		return nil
	}
	// Create a new JSON object with an array containing the methodABI object
	methodABIArray := []map[string]interface{}{methodABI}

	// Convert the methodABI map to a JSON string
	methodABIJSON, err := json.Marshal(methodABIArray)
	if err != nil {
		log.Println("Failed to marshal methodABI to JSON")
		return nil
	}

	token := token.NewToken(tokenAddress, math.Pow10(int(decimals)))
	balanceChan, errChan := token.Balance(
		ctx,
		client,
		name,
		string(methodABIJSON),
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
