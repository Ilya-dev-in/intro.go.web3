package main

import (
	"fmt"
	web "web3httpserver/server"

	"github.com/golobby/container/v3"
)

type Clienter interface {
	connect()
}

type EhtClient struct {
}

func newClient(c Clienter) *EhtClient {
	return &EhtClient{}
}

func (ec *EhtClient) connect() {
	fmt.Println("connected")
}

func main() {
	BindServices()

	var serverExit = make(chan int)
	go web.StartServerListener(serverExit)

	<-serverExit
}

func BindServices() {
	container.NamedSingleton("eth", func() Clienter {
		return newClient(&EhtClient{})
	})
}
