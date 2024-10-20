package blockchain

import (
	"context"
	"fmt"
	"log"

	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/rpc"
)

type RpcType int64

const (
	Btc RpcType = 0
	Eth RpcType = 1
)

func NewRpcClient(rpcType RpcType) {
	c := client.NewClient(rpc.MainnetRPCEndpoint)
	resp, err := c.GetVersion(context.TODO())
	if err != nil {
		log.Fatalf("failed to version info, err: %v", err)
	}

	fmt.Println("version", resp.SolanaCore)
}
