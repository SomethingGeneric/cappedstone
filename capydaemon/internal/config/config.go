package config

type Config struct {
    TLSCertFile string `json:"tls_cert_file"`
    TLSKeyFile  string `json:"tls_key_file"`
    ServerAddr  string `json:"server_addr"`
    Port        int    `json:"port"`
}