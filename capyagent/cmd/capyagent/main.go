package main

import (
	"flag"
	"log"
	"net"
	"time"

	"capyagent/internal/client"
)

func main() {
	configPath := flag.String("config", "config.json", "path to configuration file")
	flag.Parse()

	// Load configuration
	cfg, err := client.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	interval := time.Duration(cfg.PollIntervalMinutes) * time.Minute
	log.Printf("Starting client; contacting %s every %s", cfg.ServerAddress, interval)

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	runOnce := func() {
		conn, err := net.Dial("tcp", cfg.ServerAddress)
		if err != nil {
			log.Printf("Error connecting to server: %v", err)
			return
		}
		defer conn.Close()

		// Execute client logic
		if err := client.Execute(conn); err != nil {
			log.Printf("Error executing client logic: %v", err)
			return
		}

		log.Println("Client executed successfully")
	}

	// Execute immediately once before starting interval loop
	runOnce()

	for range ticker.C {
		runOnce()
	}
}
