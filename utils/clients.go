package utils

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Clients struct {
	Ctx   context.Context
	cache map[string]*ethclient.Client
}

func NewClients() *Clients {
	ctx := context.Background()
	return &Clients{
		Ctx:   ctx,
		cache: make(map[string]*ethclient.Client),
	}
}

func (s *Clients) GetClient(network string) (chan *ethclient.Client, chan error) {
	clientChan := make(chan *ethclient.Client)
	errChan := make(chan error)

	go func() {
		if client, ok := s.cache[network]; ok {
			clientChan <- client
			return
		}

		rpc := GetNetwork(network).Rpc
		if len(rpc) == 0 {
			err := fmt.Sprintf("RPC not found for network %s\n", network)
			log.Printf(err)
			errChan <- errors.New(err)
			return
		}

		url, ok := rpc[0].(string)
		if !ok {
			err := fmt.Sprintf("Url not found for network %s\n", network)
			log.Printf(err)
			errChan <- errors.New(err)
			return
		}

		client, err := ethclient.DialContext(s.Ctx, url)
		if err != nil {
			err := fmt.Sprintf("Error creating client for network %s, %s\n", network, err)
			log.Printf(err)
			errChan <- errors.New(err)
			return
		}

		s.cache[network] = client
		clientChan <- client
	}()

	return clientChan, errChan
}
