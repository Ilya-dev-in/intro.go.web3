package rpc_dto

type Asset struct {
	Interface string `json:"interface"`
	Content   *struct {
		Metadata *struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"metadata"`

		TokenInfo *struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Supply   uint32 `json:"supply"`
			Decimals int    `json:"decimals"`

			PriceInfo *struct {
				Price uint32 `json:"price_per_token"`
			} `json:"price_info"`
		} `json:"token_info"`
	} `json:"content"`
}
