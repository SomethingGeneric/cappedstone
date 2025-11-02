package config

type Config struct {
	ServerAddress string `json:"server_address"`
	Timeout       int    `json:"timeout"`
}
