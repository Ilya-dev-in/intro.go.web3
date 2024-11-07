package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"
	"web3httpserver/app"
	web "web3httpserver/app/server"
	"web3httpserver/app/services"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	services.BindServices()
	addLog()
	var serverExit = make(chan int)
	go web.StartServerListener(serverExit)

	<-serverExit
}

func addLog() {
	cfg := app.ResolveAppConfig()
	lvl, lvlErr := zerolog.ParseLevel(cfg.Logger.LogLevel)
	if lvlErr != nil {
		fmt.Println("Log level is not defined - set by default debug")
		lvl = zerolog.DebugLevel
	}
	zerolog.SetGlobalLevel(lvl)

	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		fmt.Println()
		return filepath.Base(file) + ":" + strconv.Itoa(line)
	}
	os.MkdirAll("logs", os.ModePerm)
	file, err := os.OpenFile(fmt.Sprintf("logs/%s.log", time.Now().Format("02.01.2006")), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	logger := zerolog.New(file).With().Timestamp().Logger()
	log.Logger = logger
}
