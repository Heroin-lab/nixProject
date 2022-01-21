package appserver

import (
	"github.com/Heroin-lab/nixProject/repositories/database"
)

type Config struct {
	BindAddress        string `toml:"bind_address"`
	LogLevel           string `toml:"log_level"`
	AccessSecretStr    string `toml:"accessSecret"`
	RefreshSecretStr   string `toml:"refreshSecret"`
	AccessLifetimeMin  int    `toml:"accessLifetimeMinutes"`
	RefreshLifetimeMin int    `toml:"refreshLifetimeMinutes"`
	Storage            *database.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddress:        ":7777",
		LogLevel:           "debug",
		AccessSecretStr:    "access_secret_k",
		RefreshSecretStr:   "refresh_secret_k",
		AccessLifetimeMin:  5,
		RefreshLifetimeMin: 60,
		Storage:            database.NewConfig(),
	}
}
