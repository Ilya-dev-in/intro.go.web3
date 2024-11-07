package rpc_dto

type RpcResult[T any] struct {
	Result  T      `json:"result"`
	JsonRpc string `json:"result"`
}
