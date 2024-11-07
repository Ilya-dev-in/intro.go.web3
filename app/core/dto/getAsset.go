package rpc_dto

type Asset struct {
	Interface string `json:"interface"`
	Name      string `json:"content.metadata.name"`
	Symbol    string `json:"content.metadata.symbol"`
}
