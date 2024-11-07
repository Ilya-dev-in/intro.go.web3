package services

import (
	app "web3httpserver/app"
	blockchain "web3httpserver/app/core"

	"github.com/golobby/container/v3"
)

func BindServices() {
	container.Singleton(func() app.AppConfig {
		return app.InitEnv()
	})
	container.NamedSingleton("Solana", func() blockchain.Clienter {
		return blockchain.NewSolanaClient()
	})
}
