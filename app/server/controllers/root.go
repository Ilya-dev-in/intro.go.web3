package web

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	httpHandler "web3httpserver/app/server/utils"

	"github.com/blocto/solana-go-sdk/client"
	w3client "github.com/blocto/solana-go-sdk/client"
	rpc "github.com/blocto/solana-go-sdk/rpc"
)

func GetRoot(w http.ResponseWriter, r *http.Request, _ *any) (any, error) {
	return struct{ Name string }{Name: "Stay connected"}, nil
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
	c := w3client.NewClient(rpc.MainnetRPCEndpoint)
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

func GetAccountSignatures(w http.ResponseWriter, r *http.Request, before *httpHandler.JsonStringValue) (any, error) {
	c := w3client.NewClient("https://mainnet.helius-rpc.com/?api-key=193c93fe-5752-487d-9749-8130b7b13c47")
	signatures, err := c.GetSignaturesForAddressWithConfig(
		r.Context(),
		"5C5oKvXWWA9T2GtG5rHsp4xQoF4G93TqwbWJ91Po3Ric",
		w3client.GetSignaturesForAddressConfig{
			Limit:  10,
			Before: before.Value,
		})
	if err != nil {
		log.Fatalf("failed to account signatures, err: %v", err)
		return nil, err
	}

	return signatures, nil
}

func GetTokenTransactions(w http.ResponseWriter, r *http.Request, _ *any) (any, error) {
	c := client.NewClient(rpc.MainnetRPCEndpoint)
	signatures, err := c.GetSignaturesForAddressWithConfig(
		context.TODO(),
		"5C5oKvXWWA9T2GtG5rHsp4xQoF4G93TqwbWJ91Po3Ric",
		client.GetSignaturesForAddressConfig{
			Limit: 5,
		})
	if err != nil {
		log.Fatalf("failed to account signatures, err: %v", err)
		return nil, err
	}

	transactions := make([]client.Transaction, len(signatures))

	//var wg sync.WaitGroup
	//semaphore := NewSemaphore(2)
	// for index, signature := range signatures {
	// 	//wg.Add(1)
	// 	//go func(signatureStr string) {
	// 	//semaphore.Acquire()
	// 	// wg.Done()
	// 	//defer semaphore.Release()
	// 	time.Sleep(time.Second * 5)
	// 	transaction, _ := GetSignatureTransaction(signature.Signature)
	// 	transactions[index] = *transaction
	// 	//fmt.Println(transaction?.Meta)
	// 	//transactions[index] = *transaction
	// 	//}(signature.Signature)
	// }

	//wg.Wait()

	return transactions, nil
}

type Semaphore struct {
	semaCh chan struct{}
}

func NewSemaphore(maxReq int) *Semaphore {
	return &Semaphore{
		semaCh: make(chan struct{}, maxReq),
	}
}

func (s *Semaphore) Acquire() {
	s.semaCh <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.semaCh
}

func GetSignatureTransaction(w http.ResponseWriter, r *http.Request, signature *httpHandler.JsonStringValue) (any, error) {
	c := client.NewClient("https://mainnet.helius-rpc.com/?api-key=193c93fe-5752-487d-9749-8130b7b13c47") //rpc.MainnetRPCEndpoint)
	time.Sleep(time.Second * 2)
	transaction, err := c.GetTransaction(context.TODO(), signature.Value)
	if err != nil {
		log.Fatalf("failed to get transaction info, err: %v", err)
		return nil, err
	}

	return transaction, nil
}

func GetAssetByMint(w http.ResponseWriter, r *http.Request, mint *httpHandler.JsonStringValue) (any, error) {
	c := client.NewClient("https://mainnet.helius-rpc.com/?api-key=193c93fe-5752-487d-9749-8130b7b13c47") //rpc.MainnetRPCEndpoint)
	asset, _ := GetAsset(c, mint.Value, r.Context())

	return asset, nil
}

func GetAsset(c *client.Client, address string, ctx context.Context) (any, error) {
	asset, err := call[any](&c.RpcClient, ctx, "getAsset", address)
	if err != nil {
		return nil, err
	}

	return asset, err
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

/*func GetTransactions(doneCh chan struct{}, address string) chan rpc.GetSignaturesForAddress {
	c := client.NewClient(rpc.MainnetRPCEndpoint)
	outputCh := make(chan rpc.GetSignaturesForAddress)
	signatures, err := c.GetSignaturesForAddressWithConfig(
		context.TODO(),
		"5C5oKvXWWA9T2GtG5rHsp4xQoF4G93TqwbWJ91Po3Ric",
		client.GetSignaturesForAddressConfig{
			Limit: 100,
		})
	if err != nil {
		<-doneCh
		return nil
	}

	go func() {
		defer close(outputCh)

		for _, signature := range signatures {
			select {
			case <-doneCh:
				return
			case outputCh <- signature:
			}
		}
	}()

	return outputCh
}*/
