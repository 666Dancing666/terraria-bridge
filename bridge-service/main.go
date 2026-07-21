package main

import (
"log"
"terraria-bridge/config"
"terraria-bridge/logger"
"terraria-bridge/mapping"
"terraria-bridge/mcproto"
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
mapping.InitSounds()
mapping.InitRecipes()
mapping.InitHooks()

log.Printf("Mappings: %d tiles, %d items, %d walls, %d entities, %d sounds, %d recipes, %d hooks",
len(mapping.Tiles), len(mapping.Items),
len(mapping.Walls), len(mapping.Entities),
len(mapping.Sounds), len(mapping.Recipes), len(mapping.Hooks))
log.Println("Log path:", cfg.LogPath)

go mcproto.StartMCProxy("25565")
server.Start(cfg)
}
