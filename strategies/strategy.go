package strategies

import (
	"log"
	"math/big"

	"github.com/This-Is-Prince/strategiesGo/utils"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	ERC20_BALANCE_OF   = "erc20-balance-of"
	ERC20_WITH_BALANCE = "erc20-with-balance"
	WHITELIST          = "whitelist"
	TICKET             = "ticket"

	ERC_721                 = "erc721"
	ERC_721_WITH_MULTIPLIER = "erc721-with-multiplier"
	ERC_721_ENUMERABLE      = "erc721-enumerable"
	ERC_721_WITH_TOKENID    = "erc721-with-tokenid"

	ETH_BALANCE         = "eth-balance"
	ETH_WITH_BALANCE    = "eth-with-balance"
	CONTRACT_CALL       = "contract-call"
	MULTICHAIN          = "multichain"
	ERC_1155_BALANCE_OF = "erc1155-balance-of"
	ENS_DOMAIN_OWNED    = "ens-domain-owned"
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
	case ENS_DOMAIN_OWNED:
		s.Params["network"] = s.Network
		return ENSDomainOwned(clients.Ctx, address, s.Params, nil)
	}

	// These strategy require clients
	switch s.Name {
	case MULTICHAIN:
		return Multichain(clients.Ctx, address, s.Params, clients, nil)
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
	case ERC20_WITH_BALANCE:
		return ERC20WithBalance(clients.Ctx, address, s.Params, client, nil)

	case ERC_721:
		return ERC721(clients.Ctx, address, s.Params, client, nil)
	case ERC_721_WITH_MULTIPLIER:
		return ERC721WithMultiplier(clients.Ctx, address, s.Params, client, nil)
	case ERC_721_ENUMERABLE:
		return ERC721Enumerable(clients.Ctx, address, s.Params, client, nil)
	case ERC_721_WITH_TOKENID:
		return ERC721WithTokenId(clients.Ctx, address, s.Params, client, nil)

	case ERC_1155_BALANCE_OF:
		return ERC1155BalanceOf(clients.Ctx, address, s.Params, client, nil)
	case ETH_BALANCE:
		s.Params["address"] = utils.GetNetwork(s.Network).Multicall
		return EthBalance(clients.Ctx, address, s.Params, client, nil)
	case ETH_WITH_BALANCE:
		s.Params["address"] = utils.GetNetwork(s.Network).Multicall
		return EthWithBalance(clients.Ctx, address, s.Params, client, nil)
	case CONTRACT_CALL:
		return ContractCall(clients.Ctx, address, s.Params, client, nil)
	}
	return nil
}
