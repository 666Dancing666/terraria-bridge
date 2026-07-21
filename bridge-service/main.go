package main

import (
    "terraria-bridge/config"
    "terraria-bridge/server"
)

func main() {
    cfg := config.Load()
    server.Start(cfg)
}
