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
	TICKET           = "ticket"
	ERC_721          = "erc721"
	ETH_BALANCE      = "eth-balance"
)

type Strategy struct {
	Name    string                 `json:"name"`
	Network string                 `json:"network"`
	Params  map[string]interface{} `json:"params"`
}

func (s *Strategy) Score(clients *utils.Clients, address string) *big.Float {
	// These strategy don't require client
	switch s.Name {
	case WHITELIST:
		return Whitelist(clients.Ctx, address, s.Params)
	case TICKET:
		return Ticket(clients.Ctx, address, s.Params)
	}

	var client *ethclient.Client
	clientChan, errChan := clients.GetClient(s.Network)

	select {
	case c := <-clientChan:
		client = c
	case err := <-errChan:
		log.Println(err)
		return nil
	}

	// These strategy require client
	switch s.Name {
	case ERC20_BALANCE_OF:
		return ERC20BalanceOf(clients.Ctx, address, s.Params, client, nil)
	case ERC_721:
		return ERC721(clients.Ctx, address, s.Params, client, nil)
	case ETH_BALANCE:
		return EthBalance(clients.Ctx, address, map[string]interface{}{
			"address": utils.GetNetwork(s.Network).Multicall,
		}, client, nil)
	}
	return nil
}
