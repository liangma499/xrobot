package walletsdk

import (
	"context"

	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/rpc"
)

type Client struct {
	client *client.Client
}

const url = "https://solana-mainnet.g.alchemy.com/v2/9xZ73WTfCbCpKqm0FQ9DMWCMP7v8Q499"

func NewClient() *Client {

	return &Client{
		client: client.NewClient(url),
	}
}
func (c *Client) GetBalance(base58Addr string) (uint64, error) {
	return c.client.GetBalance(
		context.TODO(),
		base58Addr,
	)
}

func (c *Client) GetBalanceWithConfig(base58Addr string, commitment rpc.Commitment) (uint64, error) {
	return c.client.GetBalanceWithConfig(
		context.TODO(),
		base58Addr,
		client.GetBalanceConfig{
			Commitment: commitment,
		},
	)
}
func (c *Client) GetBalanceFull(base58Addr string) (rpc.JsonRpcResponse[rpc.ValueWithContext[uint64]], error) {
	return c.client.RpcClient.GetBalance(
		context.TODO(),
		base58Addr,
	)
}

func (c *Client) GetTransaction(txrobot string) (*client.Transaction, error) {
	return c.client.GetTransaction(
		context.TODO(),
		txrobot,
	)
}
func (c *Client) GetSlot() (uint64, error) {
	return c.client.GetSlot(
		context.TODO(),
	)
}

func (c *Client) GetBlock(slot uint64) (*client.Block, error) {
	return c.client.GetBlock(
		context.TODO(),
		slot,
	)
}
