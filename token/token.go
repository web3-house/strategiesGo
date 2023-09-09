package token

import (
	"context"
	"errors"
	"log"
	"math/big"
	"strings"

	abiUtils "github.com/This-Is-Prince/strategiesGo/utils/abi"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Token struct {
	address  common.Address
	decimals float64
}

func NewToken(address string, decimals float64) *Token {
	return &Token{
		address:  common.HexToAddress(address),
		decimals: decimals,
	}
}

func (t *Token) Format(value *big.Int) *big.Float {
	if t.decimals != 0 {
		return new(big.Float).Quo(new(big.Float).SetInt(value), new(big.Float).SetFloat64(t.decimals))
	}
	return new(big.Float).SetInt(value)
}

func (t *Token) Balance(ctx context.Context, client *ethclient.Client, name string, address common.Address, blockNumber *big.Int) (chan *big.Int, chan error) {
	balanceChan := make(chan *big.Int)
	errChan := make(chan error)

	go func() {
		// Create a new instance of the contract
		contractAbi, err := abi.JSON(strings.NewReader(abiUtils.GetABI(name)))
		if err != nil {
			log.Println("Failed to parse contract ABI")
			errChan <- err
			return
		}

		// Call the balanceOf function
		callData, err := contractAbi.Pack(name, address)
		if err != nil {
			log.Println("Failed to pack function call data")
			errChan <- err
			return
		}

		msg := ethereum.CallMsg{
			To:   &t.address,
			Data: callData,
		}

		// Call the contract function
		contractResult, err := client.CallContract(ctx, msg, blockNumber)
		if err != nil {
			log.Println("Failed to call contract function")
			errChan <- err
			return
		}

		// Decode the result
		balanceResult, err := contractAbi.Unpack(name, contractResult)
		if err != nil {
			log.Println("Failed to unpack function result")
			errChan <- err
			return
		}

		if len(balanceResult) > 0 {
			bal, ok := balanceResult[0].(*big.Int)
			if ok {
				balanceChan <- bal
				return
			}
		}
		errChan <- errors.New("Balance is not found")
	}()

	return balanceChan, errChan
}
