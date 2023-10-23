package main

import web "web3httpserver/server"

func main() {
	var serverExit = make(chan int)
	go web.StartServerListener(serverExit)

	<-serverExit
}
