package web

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"web3httpserver/server/controllers"
)

func StartServerListener(exit chan int) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/root", controllers.GetRoot)
	mux.HandleFunc("/api/accountBalance", controllers.GetAccountBalance)
	mux.HandleFunc("/api/getTokens", controllers.GetTokens)
	mux.Handle("/api/test", BaseHttpHandler{controllers.GetTestToken})

	err := http.ListenAndServe(":3333", mux)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
		<-exit
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		<-exit
		os.Exit(1)
	}

	defer func() {
		fmt.Println("fdgd")
	}()
}

type BaseHttpResponse struct {
	Result  any   `json:"result"`
	Success bool  `json:"success"`
	Error   error `json:"error"`
}

type BaseHttpHandler struct {
	handleFn func(http.ResponseWriter, *http.Request) (any, error)
}

func (handler BaseHttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	result, err := handler.handleFn(w, r)
	response := BaseHttpResponse{
		Success: err == nil,
		Result:  result,
		Error:   err,
	}

	jsonStr, _ := json.Marshal(response)
	w.Write(jsonStr)
}
