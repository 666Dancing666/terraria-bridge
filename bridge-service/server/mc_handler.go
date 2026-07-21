package server

import (
"log"
"terraria-bridge/converter"
"terraria-bridge/filter"
"terraria-bridge/logger"
"terraria-bridge/mapping"
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

switch msg.Type {
case "request_recipes":
recipes := make([]map[string]interface{}, 0)
for key, recipe := range mapping.Recipes {
ingredients := make([]map[string]interface{}, 0)
for id, count := range recipe.Ingredients {
item := mapping.Items[id]
ingredients = append(ingredients, map[string]interface{}{
"item_id":  id,
"item_name": item.TRName,
"mc_name":  item.MCName,
"count":    count,
})
}
resultItem := mapping.Items[recipe.ResultItem]
recipes = append(recipes, map[string]interface{}{
"key":         key,
"result_id":   recipe.ResultItem,
"result_name": resultItem.TRName,
"result_mc":   resultItem.MCName,
"result_count": recipe.ResultCount,
"station":     recipe.Station,
"ingredients": ingredients,
})
}
conn.WriteJSON(protocol.Message{
Type: "recipe_list",
Payload: map[string]interface{}{
"recipes": recipes,
},
})
continue

case "craft_item":
recipeKey, _ := msg.Payload["recipe"].(string)
recipe := mapping.CheckRecipe(recipeKey)
if recipe != nil {
craftMsg := protocol.Message{
Type: "craft_request",
Payload: map[string]interface{}{
"player_id":   playerID,
"result_id":   recipe.ResultItem,
"result_count": recipe.ResultCount,
"station":     recipe.Station,
"ingredients": recipe.Ingredients,
},
}
tsConn := session.Manager.GetTShock()
if tsConn != nil {
tsConn.WriteJSON(craftMsg)
}
conn.WriteJSON(protocol.Message{
Type: "craft_result",
Payload: map[string]interface{}{
"success": true,
"recipe":  recipeKey,
},
})
} else {
conn.WriteJSON(protocol.Message{
Type: "craft_result",
Payload: map[string]interface{}{
"success": false,
"error":   "recipe not found",
},
})
}
continue
}

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
