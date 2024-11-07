package web

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
	app "web3httpserver/app"
	controllers "web3httpserver/app/server/controllers"
	httpHandler "web3httpserver/app/server/utils"
)

func StartServerListener(exit chan int) {
	mux := http.NewServeMux()
	mux.Handle("/api/root", httpHandler.BaseHttpHandler[any](controllers.GetRoot))
	mux.HandleFunc("/api/accountBalance", controllers.GetAccountBalance)
	mux.Handle("/api/getTokens", httpHandler.BaseHttpHandler[httpHandler.BaseHttpRequestPagination](controllers.GetTokens))
	mux.Handle("/api/getTokensTransactions", httpHandler.BaseHttpHandler[any](controllers.GetTokenTransactions))

	mux.Handle("/api/getAccountSignatures", httpHandler.BaseHttpHandler[httpHandler.JsonStringValue](controllers.GetAccountSignatures))

	mux.Handle("/api/GetSignatureTransaction", httpHandler.BaseHttpHandler[httpHandler.JsonStringValue](controllers.GetSignatureTransaction))

	mux.HandleFunc("/api/test", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		handleConnection(w)
	})

	mux.Handle("/api/GetAssetByMint", httpHandler.BaseHttpHandler[httpHandler.JsonStringValue](controllers.GetAssetByMint))

	err := http.ListenAndServe(fmt.Sprintf(":%s", app.ResolveAppConfig().App.Port), mux)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
		<-exit
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		<-exit
		os.Exit(1)
	}
}

func handleConnection(w http.ResponseWriter) {
	time.Sleep(10 * time.Second)
	w.Header().Add("Connection", "keep-alive")
	w.WriteHeader(200)
}
