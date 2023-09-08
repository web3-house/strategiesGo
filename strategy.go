package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Network struct {
	Name      string        `json:"name"`
	Key       string        `json:"key"`
	ChainId   int64         `json:"chainId"`
	Network   string        `json:"network"`
	Multicall string        `json:"multicall"`
	Rpc       []interface{} `json:"rpc"`
}

type Strategy struct {
	client   *ethclient.Client
	ctx      context.Context
	networks map[string]Network
}

func NewStrategy(network string) *Strategy {
	ctx := context.Background()
	client, err := ethclient.DialContext(ctx, url)
	if err != nil {
		log.Printf("Error creating client for network %s, %s", network, err)
	}
	data, err := ioutil.ReadFile("networks.json")
	if err != nil {
		log.Printf("Error while loading networks.json %s", err)
	}

	// Parse JSON into Go struct
	var networksJSON map[string]Network
	err = json.Unmarshal(data, &networksJSON)
	if err != nil {
		log.Printf("Error while unmarshal networks.json %s", err)
	}
	return &Strategy{
		client:   client,
		ctx:      ctx,
		networks: networksJSON,
	}
}
