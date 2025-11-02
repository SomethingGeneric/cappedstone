package main

import (
	"fmt"
	"log"

	"capydaemon/internal/config"
	"capydaemon/internal/server"
	tlsutil "capydaemon/internal/tls"
)

func main() {
	// Load server configuration
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	if err := tlsutil.EnsureCertsExist(cfg.TLSCertFile, cfg.TLSKeyFile); err != nil {
		log.Fatalf("TLS certificate files missing: %v", err)
	}

	tlsConfig, err := tlsutil.LoadTLSConfig(cfg.TLSCertFile, cfg.TLSKeyFile)
	if err != nil {
		log.Fatalf("Failed to load TLS configuration: %v", err)
	}

	address := fmt.Sprintf("%s:%d", cfg.ServerAddr, cfg.Port)
	serverConfig := &server.Config{Address: address, TLSConfig: tlsConfig}
	daemon, err := server.NewServer(serverConfig)
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	daemon.Start()
}
