package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Network struct {
	Name      string        `json:"name"`
	Key       string        `json:"key"`
	ChainId   int64         `json:"chainId"`
	Network   string        `json:"network"`
	Multicall string        `json:"multicall"`
	Rpc       []interface{} `json:"rpc"`
}

var networks = map[string]Network{}

func init() {
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
	networks = networksJSON
}

func GetNetwork(key string) Network {
	network, ok := networks[key]
	if ok {
		return network
	}
	return Network{}
}
