package strategies

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/This-Is-Prince/strategiesGo/utils"
)

func Multichain(ctx context.Context, address string, params map[string]interface{}, clients *utils.Clients, blockNumber *big.Int) *big.Float {

	strategiesValue, ok := params["strategies"]
	if !ok {
		return nil
	}
	// Convert the interface{} to JSON bytes
	strategiesBytes, err := json.Marshal(strategiesValue)
	if err != nil {
		fmt.Println("JSON marshaling error:", err)
		return nil
	}

	// Unmarshal JSON bytes into a struct
	var strategies []Strategy
	if err := json.Unmarshal(strategiesBytes, &strategies); err != nil {
		fmt.Println("JSON un marshaling error:", err)
		return nil
	}

	score := big.NewFloat(0)
	for _, strategy := range strategies {
		score = score.Add(score, strategy.Score(clients, address))
	}

	return score
}
