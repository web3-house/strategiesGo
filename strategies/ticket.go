package strategies

import (
	"context"
	"math/big"
)

func Ticket(ctx context.Context, address string, params map[string]interface{}) *big.Float {
	value, ok := params["value"]
	if !ok {
		return big.NewFloat(1)
	}

	v, ok := value.(float64)
	if !ok {
		return big.NewFloat(1)
	}

	return big.NewFloat(v)
}
