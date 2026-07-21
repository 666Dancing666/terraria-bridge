package main

import (
"log"
"terraria-bridge/config"
"terraria-bridge/logger"
"terraria-bridge/mapping"
"terraria-bridge/server"
)

func main() {
cfg := config.Load()

if err := logger.Init(cfg.LogPath); err != nil {
log.Fatalf("Failed to init logger: %v", err)
}
defer logger.Default.Close()

mapping.InitTiles()
mapping.InitItems()
mapping.InitEntities()

log.Printf("Mappings loaded: %d tiles, %d items, %d walls, %d entities",
len(mapping.Tiles), len(mapping.Items), len(mapping.Walls), len(mapping.Entities))
log.Println("Log path:", cfg.LogPath)
server.Start(cfg)
}
