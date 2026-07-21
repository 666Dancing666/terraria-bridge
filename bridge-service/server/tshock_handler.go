package server

import (
"log"
"terraria-bridge/converter"
"terraria-bridge/filter"
"terraria-bridge/logger"
"terraria-bridge/protocol"
"terraria-bridge/session"

"github.com/gorilla/websocket"
)

func handleTShock(conn *websocket.Conn) {
log.Println("TShock connected")
session.Manager.SetTShock(conn)

for _, msg := range session.Manager.TSBuffer.PopAll() {
conn.WriteJSON(msg)
}

for {
var msg protocol.Message
err := conn.ReadJSON(&msg)
if err != nil {
log.Printf("TShock disconnected: %v", err)
session.Manager.SetTShock(nil)
break
}

if !filter.Default.PassIn(msg.Type) {
filter.Default.RecordBlocked(msg.Type)
continue
}

logger.Default.Log("TSHOCK->MC", msg.Type, msg.Payload)

var toSend protocol.Message
switch msg.Type {
case "world_snapshot":
toSend = converter.ConvertWorldSnapshot(msg)
case "entity_update":
toSend = converter.ConvertEntityUpdate(msg)
case "tile_update", "wall_update", "liquid_update":
toSend = converter.ConvertTileUpdate(msg)
case "player_move":
toSend = converter.ConvertPlayerMove(msg)
case "game_event":
toSend = converter.ConvertEvent(msg)
case "time_sync":
toSend = converter.ConvertTimeSync(msg)
case "weather_sync":
toSend = converter.ConvertWeatherSync(msg)
case "chat_message":
toSend = converter.ConvertChatTRToMC(msg)
default:
toSend = msg
}

for _, ps := range session.Manager.GetAllPlayers() {
if ps.MCConn != nil {
if err := ps.MCConn.WriteJSON(toSend); err != nil {
ps.MCBuffer.Push(toSend)
}
}
}
}
}
