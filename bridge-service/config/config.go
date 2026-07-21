package config

import "os"

type Config struct {
Port    string
LogPath string
}

func Load() *Config {
port := os.Getenv("BRIDGE_PORT")
if port == "" {
port = "8080"
}
logPath := os.Getenv("BRIDGE_LOG_PATH")
if logPath == "" {
logPath = "logs"
}
return &Config{Port: ":" + port, LogPath: logPath}
}
