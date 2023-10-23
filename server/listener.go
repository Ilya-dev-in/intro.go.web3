package web

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"web3httpserver/server/controllers"
)

func StartServerListener(exit chan int) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/root", controllers.GetRoot)

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
