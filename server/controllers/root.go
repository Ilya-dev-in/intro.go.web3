package web

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	httpHandler "web3httpserver/server/utils"

	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/rpc"
)

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

func GetTokens(w http.ResponseWriter, r *http.Request, t *httpHandler.BaseHttpRequestPagination) (any, error) {
	c := client.NewClient(rpc.MainnetRPCEndpoint)
	tokenAccounts, err := c.GetTokenAccountsByOwnerByProgram(
		context.TODO(),
		"5C5oKvXWWA9T2GtG5rHsp4xQoF4G93TqwbWJ91Po3Ric",
		"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA")
	if err != nil {
		log.Fatalf("failed to version info, err: %v", err)
		return nil, err
	}

	return tokenAccounts, nil
}

func GetTokenTransactions(w http.ResponseWriter, r *http.Request, _ *any) (any, error) {
	c := client.NewClient(rpc.MainnetRPCEndpoint)
	tokenAccounts, err := c.GetSignaturesForAddressWithConfig(
		context.TODO(),
		"5C5oKvXWWA9T2GtG5rHsp4xQoF4G93TqwbWJ91Po3Ric",
		client.GetSignaturesForAddressConfig{
			Limit: 1,
		})
	if err != nil {
		log.Fatalf("failed to version info, err: %v", err)
		return nil, err
	}

	transaction, tErr := c.GetTransaction(context.TODO(), tokenAccounts[0].Signature)
	if tErr != nil {
		log.Fatalf("failed to get transaction info, err: %v", err)
		return nil, err
	}

	fmt.Println(transaction.Meta)

	return tokenAccounts, nil
}
