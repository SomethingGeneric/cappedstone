package main

import (
    "log"
    "net"
    "os"

    "capyagent/internal/client"
)

func main() {
    // Load configuration
    cfg, err := client.LoadConfig("config.json")
    if err != nil {
        log.Fatalf("Error loading config: %v", err)
    }

    // Connect to the server
    conn, err := net.Dial("tcp", cfg.ServerAddress)
    if err != nil {
        log.Fatalf("Error connecting to server: %v", err)
    }
    defer conn.Close()

    // Execute client logic
    if err := client.Execute(conn); err != nil {
        log.Fatalf("Error executing client logic: %v", err)
    }

    log.Println("Client executed successfully")
}