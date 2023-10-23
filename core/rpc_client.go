package blockchain

import (
	"fmt"

	walletconnect "github.com/civet148/walletconnect"
	bitcoinrpc "github.com/wenweih/bitcoin-rpc-golang"
)

type RpcType int64

const (
	Btc RpcType = 0
	Eth RpcType = 1
)

func NewRpcClient(rpcType RpcType) {
	btcClient, err := bitcoinrpc.New(&bitcoinrpc.ConnConfig{
		Host:       "190.2.146.90:8333",
		User:       "user",
		Pass:       "pass",
		DisableTLS: true,
	})

	walletConn, err := walletconnect.NewDApp("sdf")
	if err != nil {
		fmt.Println(err)
	}

}
