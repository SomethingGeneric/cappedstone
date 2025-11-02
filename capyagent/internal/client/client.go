package client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os"
	"time"
)

type Client struct {
	serverAddr string
	tlsConfig  *tls.Config
}

type Request struct {
	Command string `json:"command"`
}

type Response struct {
	Result string `json:"result"`
}

type Config struct {
	ServerAddress       string `json:"server_address"`
	PollIntervalMinutes int    `json:"poll_interval_minutes"`
}

func NewClient(serverAddr string, tlsConfig *tls.Config) *Client {
	return &Client{
		serverAddr: serverAddr,
		tlsConfig:  tlsConfig,
	}
}

func LoadConfig(path string) (*Config, error) {
	const defaultPollIntervalMinutes = 5

	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			// default configuration if no config file is present
			return &Config{
				ServerAddress:       "127.0.0.1:8443",
				PollIntervalMinutes: defaultPollIntervalMinutes,
			}, nil
		}
		return nil, fmt.Errorf("unable to read config: %w", err)
	}

	var cfg Config
	if err := json.Unmarshal(bytes.TrimSpace(data), &cfg); err != nil {
		return nil, fmt.Errorf("unable to parse config: %w", err)
	}
	if cfg.ServerAddress == "" {
		return nil, fmt.Errorf("config missing server_address")
	}
	if cfg.PollIntervalMinutes <= 0 {
		cfg.PollIntervalMinutes = defaultPollIntervalMinutes
	}
	return &cfg, nil
}

func (c *Client) SendRequest(command string) (string, error) {
	conn, err := tls.Dial("tcp", c.serverAddr, c.tlsConfig)
	if err != nil {
		return "", fmt.Errorf("failed to connect to server: %w", err)
	}
	defer conn.Close()

	req := Request{Command: command}
	reqData, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	_, err = conn.Write(reqData)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}

	responseData := make([]byte, 1024)
	conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	n, err := conn.Read(responseData)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	var response Response
	if err := json.Unmarshal(bytes.TrimSpace(responseData[:n]), &response); err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return response.Result, nil
}

func Execute(conn net.Conn) error {
	if conn == nil {
		return fmt.Errorf("connection is nil")
	}
	// Future protocol handling would live here.
	return nil
}
