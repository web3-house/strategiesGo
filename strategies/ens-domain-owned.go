package strategies

import (
	"context"
	"math/big"

	ensUtils "github.com/This-Is-Prince/strategiesGo/utils/ens"
)

func ENSDomainOwned(ctx context.Context, address string, params map[string]interface{}, blockNumber *big.Int) *big.Float {
	domainValue, ok := params["domain"]
	if !ok {
		return nil
	}
	domain, ok := domainValue.(string)
	if !ok {
		return nil
	}
	networkValue, ok := params["network"]
	if !ok {
		return nil
	}
	network, ok := networkValue.(string)
	if !ok {
		return nil
	}

	args := ensUtils.Args{
		Network: network,
		Query:   ensUtils.GET_DOMAINS_WITH_DOMAIN_NAME_AND_GET_SUBDOMAINS_WITH_OWNER_ADDRESSES,
		Variables: map[string]interface{}{
			"domain_name": domain,
			"owner_in":    []string{address},
		},
	}

	data := struct {
		Domains []ensUtils.Domain `json:"domains"`
	}{}

	err := ensUtils.FetchENS(args, &data)
	if err != nil {
		return nil
	}

	if len(data.Domains) > 0 {
		return new(big.Float).SetInt64(int64(len(data.Domains[0].Subdomains)))
	}

	return nil
}
