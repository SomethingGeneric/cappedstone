package main

import (
    "log"
    "net"
    "github.com/capy/capydaemon/internal/config"
    "github.com/capy/capydaemon/internal/tls"
    "github.com/capy/capydaemon/internal/server"
)

func main() {
    // Load server configuration
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Failed to load configuration: %v", err)
    }

    // Setup TLS listener
    listener, err := tls.NewTLSListener(cfg.TLSConfig)
    if err != nil {
        log.Fatalf("Failed to create TLS listener: %v", err)
    }
    defer listener.Close()

    log.Printf("Server is listening on %s", cfg.Address)

    // Accept incoming connections
    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Printf("Failed to accept connection: %v", err)
            continue
        }

        // Handle the connection in a new goroutine
        go server.HandleConnection(conn)
    }
}