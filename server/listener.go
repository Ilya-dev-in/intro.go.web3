package web

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	controllers "web3httpserver/server/controllers"
	httpHandler "web3httpserver/server/utils"
)

func StartServerListener(exit chan int) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/root", controllers.GetRoot)
	mux.HandleFunc("/api/accountBalance", controllers.GetAccountBalance)
	mux.Handle("/api/getTokens", httpHandler.BaseHttpHandler[httpHandler.BaseHttpRequestPagination]{
		HandleFn: controllers.GetTokens,
	})
	mux.Handle("/api/getTokensTransactions", httpHandler.BaseHttpHandler[any]{
		HandleFn: controllers.GetTokenTransactions,
	})

	err := http.ListenAndServe(":3333", mux)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
		<-exit
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		<-exit
		os.Exit(1)
	}
}
