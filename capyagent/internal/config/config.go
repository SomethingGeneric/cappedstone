package config

type Config struct {
	ServerAddress       string `json:"server_address"`
	PollIntervalMinutes int    `json:"poll_interval_minutes"`
}
