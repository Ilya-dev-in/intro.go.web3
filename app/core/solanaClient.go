package blockchain

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	solClient "github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/rpc"
	"github.com/rs/zerolog/log"

	app "web3httpserver/app"
	dto "web3httpserver/app/core/dto"
)

func NewSolanaClient() Clienter {
	c := solClient.NewClient(app.ResolveAppConfig().App.DefaultRpcNodeUrl)
	var solanaClient SolanaClient
	solanaClient.Instance = c

	return solanaClient
}

type SolanaClient struct {
	Instance *solClient.Client
}

func NewClient(c Clienter) *SolanaClient {
	return &SolanaClient{}
}

func (sc SolanaClient) GetAsset(address string, ctx context.Context) (*dto.RpcResult[dto.Asset], error) {
	asset, err := call[dto.RpcResult[dto.Asset]](&sc.Instance.RpcClient, ctx, "getAsset", address)
	if err != nil {
		return nil, err
	}

	return &asset, err
}

func (sc SolanaClient) GetAssetBatch(ctx context.Context, addresses ...string) (any, error) {
	msg := "Solana GetAssetBatch not implemented"
	log.Error().Msg(msg)

	return nil, errors.New(msg)
}

func call[T any](rc *rpc.RpcClient, ctx context.Context, params ...any) (T, error) {
	var output T

	body, err := rc.Call(ctx, params...)
	if err != nil {
		return output, fmt.Errorf("rpc: call error, err: %v, body: %v", err, string(body))
	}

	err = json.Unmarshal(body, &output)
	if err != nil {
		return output, fmt.Errorf("rpc: failed to json decode body, err: %v", err)
	}

	return output, nil
}
