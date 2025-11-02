package server

import (
    "crypto/tls"
    "fmt"
    "net"
    "log"
)

type Server struct {
    listener net.Listener
    config   *Config
}

type Config struct {
    Address string
    TLSConfig *tls.Config
}

func NewServer(config *Config) (*Server, error) {
    listener, err := tls.Listen("tcp", config.Address, config.TLSConfig)
    if err != nil {
        return nil, fmt.Errorf("failed to start server: %w", err)
    }

    return &Server{
        listener: listener,
        config:   config,
    }, nil
}

func (s *Server) Start() {
    log.Printf("Server listening on %s", s.config.Address)
    for {
        conn, err := s.listener.Accept()
        if err != nil {
            log.Printf("failed to accept connection: %v", err)
            continue
        }
        go s.handleConnection(conn)
    }
}

func (s *Server) handleConnection(conn net.Conn) {
    defer conn.Close()
    log.Printf("Client connected: %s", conn.RemoteAddr().String())
    // Handle client session logic here
}