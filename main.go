package main

import (
	"context"
	"log"
	"math/big"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const url = "https://rpc.testnet.moonbeam.network"

const balanceOfABI = `
[
	{
		"constant": true,
		"inputs": [
			{
				"name": "account",
				"type": "address"
			}
		],
		"name": "balanceOf",
		"outputs": [
			{
				"name": "",
				"type": "uint256"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	}
]
`

func main() {
	ctx := context.Background()
	client, err := ethclient.DialContext(ctx, url)
	if err != nil {
		log.Fatal("Error creating client", err)
	}

	contractAddress := common.HexToAddress("0x0000000000000000000000000000000000000802")
	hexAddresses := []string{"0xc6DDE7CE20DfcCF6A2b8c998B066a3ce48911311", "0xb98Ee84a0dcECf67399d0bca3C28A105EA0268e5"}
	wg := &sync.WaitGroup{}
	for _, hexAddress := range hexAddresses {
		wg.Add(1)
		go func(hexAddress string) {
			defer wg.Done()
			balanceOf(client, common.HexToAddress(hexAddress), contractAddress)
		}(hexAddress)
	}
	wg.Wait()
}

func balanceOf(client *ethclient.Client, address, contractAddress common.Address) {
	// Create a new instance of the contract
	contractAbi, err := abi.JSON(strings.NewReader(balanceOfABI))
	if err != nil {
		log.Fatalf("Failed to parse contract ABI: %v", err)
	}

	// Call the balanceOf function
	callData, err := contractAbi.Pack("balanceOf", address)
	if err != nil {
		log.Fatalf("Failed to pack function call data: %v", err)
	}

	msg := ethereum.CallMsg{
		To:   &contractAddress,
		Data: callData,
	}

	// Call the contract function
	result, err := client.CallContract(context.Background(), msg, nil)
	if err != nil {
		log.Fatalf("Failed to call contract function: %v", err)
	}

	// Decode the result
	results, err := contractAbi.Unpack("balanceOf", result)
	if err != nil {
		log.Fatalf("Failed to unpack function result: %v", err)
	}

	for _, result := range results {
		bal, ok := result.(*big.Int)
		balanceEther := new(big.Float).Quo(new(big.Float).SetInt(bal), new(big.Float).SetFloat64(1e18))
		log.Println(ok, balanceEther.String())
	}
}

func example(client *ethclient.Client, ctx context.Context) {

	address := common.HexToAddress("0xc6DDE7CE20DfcCF6A2b8c998B066a3ce48911311")
	bal, err := client.BalanceAt(ctx, address, nil)
	if err != nil {
		log.Fatal("Error fetching balance of given address", err)
	}

	balanceEther := new(big.Float).Quo(new(big.Float).SetInt(bal), new(big.Float).SetFloat64(1e18))

	// Format and print the balance in Ether
	log.Printf("Balance of %s: %s Ether\n", address.Hex(), balanceEther.Text('f', 18))

	blockNumber, err := client.BlockNumber(ctx)
	if err != nil {
		log.Fatal("Error fetching blocknumber", err)
	}
	log.Println("Block Number:-", blockNumber)

	block, err := client.BlockByNumber(ctx, new(big.Int).SetInt64(5043583))
	if err != nil {
		log.Fatal("Error fetching block", err)
	}
	log.Println(block.Time())

	// contractAddress := common.HexToAddress("0x0000000000000000000000000000000000000802") // Replace with the contract address
	// balanceOf(client, address, contractAddress)
}
