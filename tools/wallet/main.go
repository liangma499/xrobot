package main

import (
	"context"
	"fmt"
	"log"

	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/rpc"
)

const url = "https://solana-mainnet.g.alchemy.com/v2/9xZ73WTfCbCpKqm0FQ9DMWCMP7v8Q499"

func main() {
	c := client.NewClient(url)

	// get balance
	balance, err := c.GetBalance(
		context.TODO(),
		"7rhxnLV8C77o6d8oz26AgK8x8m5ePsdeRawjqvojbjnQ",
	)
	if err != nil {
		log.Fatalf("failed to get balance, err: %v", err)
	}
	fmt.Printf("balance: %v\n", balance)

	// get balance with sepcific commitment
	balance, err = c.GetBalanceWithConfig(
		context.TODO(),
		"7rhxnLV8C77o6d8oz26AgK8x8m5ePsdeRawjqvojbjnQ",
		client.GetBalanceConfig{
			Commitment: rpc.CommitmentProcessed,
		},
	)
	if err != nil {
		log.Fatalf("failed to get balance with cfg, err: %v", err)
	}
	fmt.Printf("balance: %v\n", balance)

	// for advanced usage. fetch full rpc response
	res, err := c.RpcClient.GetBalance(
		context.TODO(),
		"7rhxnLV8C77o6d8oz26AgK8x8m5ePsdeRawjqvojbjnQ",
	)
	if err != nil {
		log.Fatalf("failed to get balance via rpc client, err: %v", err)
	}
	fmt.Printf("response: %+v\n", res)
}
