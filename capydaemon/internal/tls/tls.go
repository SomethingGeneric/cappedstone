package tls

import (
    "crypto/tls"
    "log"
    "os"
)

// LoadTLSConfig loads the TLS configuration from the specified certificate and key files.
func LoadTLSConfig(certFile, keyFile string) (*tls.Config, error) {
    cert, err := tls.LoadX509KeyPair(certFile, keyFile)
    if err != nil {
        return nil, err
    }

    return &tls.Config{
        Certificates: []tls.Certificate{cert},
    }, nil
}

// CreateTLSListener creates a TLS listener on the specified address.
func CreateTLSListener(address, certFile, keyFile string) (net.Listener, error) {
    tlsConfig, err := LoadTLSConfig(certFile, keyFile)
    if err != nil {
        return nil, err
    }

    listener, err := tls.Listen("tcp", address, tlsConfig)
    if err != nil {
        return nil, err
    }

    return listener, nil
}

// EnsureCertsExist checks if the certificate and key files exist.
func EnsureCertsExist(certFile, keyFile string) error {
    if _, err := os.Stat(certFile); os.IsNotExist(err) {
        return err
    }
    if _, err := os.Stat(keyFile); os.IsNotExist(err) {
        return err
    }
    return nil
}