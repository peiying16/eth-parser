package config

type Config struct {
	Node string
}

func NewConfig() *Config {
	return &Config{
		Node: "https://cloudflare-eth.com",
	}
}
