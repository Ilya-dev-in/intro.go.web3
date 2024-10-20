package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/rpc"
)

const LAMPORTS_PER_SOL = 1000000000

func GetRoot(w http.ResponseWriter, r *http.Request) {
	titleJson, err := json.Marshal(struct{ Name string }{Name: "Stay connected"})
	if err != nil {
		fmt.Println(err)
	}

	w.Write(titleJson)
}

func GetAccountBalance(w http.ResponseWriter, r *http.Request) {
	c := client.NewClient(rpc.MainnetRPCEndpoint)
	balance, balanceErr := c.GetBalance(context.TODO(), "5C5oKvXWWA9T2GtG5rHsp4xQoF4G93TqwbWJ91Po3Ric")
	if balanceErr != nil {
		log.Fatalf("failed to version info, err: %v", balanceErr)
	}

	n := strconv.FormatUint(balance, 10)
	jsonStr, err := json.Marshal(struct{ Balance string }{Balance: n})
	if err != nil {
		fmt.Println(err)
	}

	w.Write(jsonStr)
}

func GetTokens(w http.ResponseWriter, r *http.Request) {
	c := client.NewClient(rpc.MainnetRPCEndpoint)
	/*tokens, balanceErr := c.RpcClient.GetTokenAccountsByOwnerWithConfig(
		context.TODO(),
		"5C5oKvXWWA9T2GtG5rHsp4xQoF4G93TqwbWJ91Po3Ric",
		rpc.GetTokenAccountsByOwnerConfigFilter{
			ProgramId: "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
		},
		rpc.GetTokenAccountsByOwnerConfig{
			Encoding: rpc.AccountEncodingJsonParsed,
		},
	)

	data, ok := tokens.Result.Value[2].Account.Data.([]AccountData)
	jsonBytes, _ := json.Marshal(tokens.Result.Value[2].Account.Data)

	var accData AccountData
	err := json.Unmarshal(jsonBytes, &accData)
	if err != nil {
	}
	if !ok {
		log.Fatalf("failed to version info, err: %v", balanceErr)
	}

	fmt.Println(data)
	fmt.Println(jsonBytes)
	fmt.Println(accData)*/
	/*a, decodeErr := base64.StdEncoding.DecodeString(data[0].(string))
	if decodeErr != nil {
		log.Fatalf("failed to version info, err: %v", balanceErr)
	}
	fmt.Println(a)
	amount := binary.LittleEndian.Uint64(a[64:72])
	fmt.Println(amount)
	if balanceErr != nil {
		log.Fatalf("failed to version info, err: %v", balanceErr)
	}*/

	/*var tokenAccount = make([]TokenAccount, len(tokens.Result.Value))
	for _, v := range tokens.Result.Value {
		data, ok := v.Account.Data.([]any)
		if !ok {
			log.Fatalf("failed to version info, err: %v", balanceErr)
		}
		tokenAccount =  append(tokenAccount, TokenAccount{
			client.TokenAccount{
				token.TokenAccount{
					Mint: common.PublicKeyFromString(v.Account.Owner),
				},
			},
		})
	}*/

	tokenAccounts, err := c.GetTokenAccountsByOwnerByProgram(
		context.TODO(),
		"5C5oKvXWWA9T2GtG5rHsp4xQoF4G93TqwbWJ91Po3Ric",
		"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA")
	if err != nil {
		//log.Fatalf("failed to version info, err: %v", err)
		return
	}
	/*notEmptyTokens := []client.TokenAccount{}

	for _, token := range tokenAccounts {
		if token.Amount > 0 {
			//token.Amount = token.Amount / tokens[index].
			notEmptyTokens = append(notEmptyTokens, token)
		}
	}*/

	json, err := json.Marshal(tokenAccounts)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(json)
}

func GetTestToken(w http.ResponseWriter, r *http.Request) (any, error) {
	return "test_str", nil
}

/*
type TokenAccount struct {
	client.TokenAccount
	UIAmount uint64
}

type AccountData struct {
	Data    AccountDataInfo    `json:"parsed"`
	Program string `json:"program"`
	Space   int    `json:"space"`
}
type AccountDataInfo struct {
	Info    AccountDataInfo    `json:"info"`
}

type TokenInfo struct {
	Decimals
}
*/
