package server

import (
"log"
"terraria-bridge/converter"
"terraria-bridge/filter"
"terraria-bridge/logger"
"terraria-bridge/metrics"
"terraria-bridge/protocol"
"terraria-bridge/session"
"fmt"
"time"

"github.com/gorilla/websocket"
)

func handleMC(conn *websocket.Conn) {
playerID := fmt.Sprintf("MC-%d", time.Now().UnixNano())
log.Printf("MC client connected: %s", playerID)

ps := session.Manager.AddPlayer(playerID, conn)
defer session.Manager.RemovePlayer(playerID)

for _, msg := range ps.MCBuffer.PopAll() {
conn.WriteJSON(msg)
}

for {
var msg protocol.Message
err := conn.ReadJSON(&msg)
if err != nil {
log.Printf("MC disconnected: %s - %v", playerID, err)
break
}

metrics.Stats.RecordOut()

if !filter.Default.PassOut(msg.Type) {
filter.Default.RecordBlocked(msg.Type)
continue
}

msg.Payload["_player_id"] = playerID
logger.Default.Log("MC->TSHOCK", msg.Type, msg.Payload)

var toSend protocol.Message
if msg.Type == "chat_message" {
toSend = converter.ConvertChatMCToTR(msg)
} else {
toSend = converter.ConvertMCToTR(msg)
}

tsConn := session.Manager.GetTShock()
if tsConn != nil {
if err := tsConn.WriteJSON(toSend); err != nil {
session.Manager.TSBuffer.Push(toSend)
}
} else {
session.Manager.TSBuffer.Push(toSend)
}
}
}
