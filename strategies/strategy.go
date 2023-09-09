package strategies

import (
	"log"
	"math/big"

	"github.com/This-Is-Prince/strategiesGo/utils"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	ERC20_BALANCE_OF = "erc20-balance-of"
	WHITELIST        = "whitelist"
)

type Strategy struct {
	Name    string                 `json:"name"`
	Network string                 `json:"network"`
	Params  map[string]interface{} `json:"params"`
}

func (s *Strategy) Score(clients *utils.Clients, address string) *big.Float {
	var client *ethclient.Client
	clientChan, errChan := clients.GetClient(s.Network)

	select {
	case c := <-clientChan:
		client = c
	case err := <-errChan:
		log.Println(err)
		return nil
	}

	switch s.Name {
	case ERC20_BALANCE_OF:
		return ERC20BalanceOf(clients.Ctx, address, s.Params, client, nil)
	case WHITELIST:
		return Whitelist(clients.Ctx, address, s.Params)
	}
	return nil
}
