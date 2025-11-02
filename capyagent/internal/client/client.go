package client

import (
    "bytes"
    "crypto/tls"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net"
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

func NewClient(serverAddr string, tlsConfig *tls.Config) *Client {
    return &Client{
        serverAddr: serverAddr,
        tlsConfig:  tlsConfig,
    }
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