package appserver

type Config struct {
	BindAddress string `toml:"bind_address"`
	LogLevel    string `toml:"log_level"`
}

func NewConfig() *Config {
	return &Config{
		BindAddress: ":7777",
		LogLevel:    "debug",
	}

}
