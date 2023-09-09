package strategies

import (
	"context"
	"math/big"
	"strings"
)

func Whitelist(ctx context.Context, address string, params map[string]interface{}) *big.Float {
	addressesValue, ok := params["addresses"]
	if !ok {
		return nil
	}
	addresses, ok := addressesValue.([]interface{})
	if !ok {
		return nil
	}

	lowerCase := strings.ToLower(address)
	for _, addr := range addresses {
		addrValue, ok := addr.(string)
		if ok && lowerCase == strings.ToLower(addrValue) {
			return big.NewFloat(1)
		}
	}

	return big.NewFloat(0)
}
