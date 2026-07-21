package server

import (
"log"
"terraria-bridge/converter"
"terraria-bridge/filter"
"terraria-bridge/logger"
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
if err := conn.WriteJSON(msg); err != nil {
ps.MCBuffer.Push(msg)
break
}
}

for {
var msg protocol.Message
err := conn.ReadJSON(&msg)
if err != nil {
log.Printf("MC disconnected: %s - %v", playerID, err)
break
}

if !filter.Default.PassOut(msg.Type) {
filter.Default.RecordBlocked(msg.Type)
continue
}

msg.Payload["_player_id"] = playerID
logger.Default.Log("MC->TSHOCK", msg.Type, msg.Payload)

toSend := converter.ConvertMCToTR(msg)

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
