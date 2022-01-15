package appserver

import (
	"github.com/Heroin-lab/nixProject/repositories/database"
)

type Config struct {
	BindAddress        string `toml:"bind_address"`
	LogLevel           string `toml:"log_level"`
	AccessSecretStr    string `toml:"accessSecret"`
	RefreshSecretStr   string `toml:"refreshSecret"`
	AccessLifetimeMin  string `toml:"accessLifetimeMinutes"`
	RefreshLifetimeMin string `toml:"refreshLifetimeMinutes"`
	Storage            *database.Config
}

//func NewConfig() *Config {
//	return &Config{
//		BindAddress: ":7777",
//		LogLevel:    "debug",
//		Storage:     database.NewConfig(),
//	}
//}
