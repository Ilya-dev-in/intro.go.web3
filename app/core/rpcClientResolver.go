package blockchain

import (
	"context"
	dto "web3httpserver/app/core/dto"

	"github.com/golobby/container/v3"
)

type Clienter interface {
	GetAsset(address string, ctx context.Context) (*dto.RpcResult[dto.Asset], error)
	GetAssetBatch(ctx context.Context, addresses ...string) (any, error)
}

type RpcType string

const (
	Btc RpcType = "BTC"
	Eth RpcType = "ETH"
	Sol RpcType = "Solana"
)

func ResolveRpcClient(t string) Clienter {
	var client Clienter
	container.NamedResolve(&client, t)

	return client
}
