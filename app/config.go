package app

import (
	"context"
	"fmt"
	"os"

	"github.com/golobby/container/v3"
	"github.com/rs/zerolog/log"
	"github.com/sethvargo/go-envconfig"
	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	App    *Server `yaml:"app"`
	Logger *Logger `yaml:"log"`
}

func ResolveAppConfig() *AppConfig {
	var cfg AppConfig
	resolveErr := container.Resolve(&cfg)
	if resolveErr != nil {
		fmt.Printf("Error to resolve AppConfig: %s", resolveErr)
		log.Error().Err(resolveErr)
		panic(resolveErr)
	}

	return &cfg
}

type Server struct {
	DefaultRpcNodeUrl string `yaml:"DEFAULT_RPC_URL", env:"DEFAULT_RPC_URL"`
	Port              string `yaml:"PORT", env:"PORT", default=3333"`
}

type Logger struct {
	LogLevel string `yaml:"LOG_LEVEL", env:"LOG_LEVEL"`
}

func InitEnv() AppConfig {
	ctx := context.Background()
	var cfg AppConfig
	readFile(&cfg)
	readEnv(ctx, &cfg)
	fmt.Printf("%+v", cfg)

	return cfg
}

func readFile(cfg *AppConfig) {
	envFile := os.Getenv("APP_ENV")
	if len(envFile) == 0 {
		envFile = ".env"
	}
	f, err := os.Open(envFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		panic(err)
	}
}

func readEnv(ctx context.Context, cfg *AppConfig) {
	err := envconfig.Process(ctx, cfg)
	if err != nil {
		panic(err)
	}
}
