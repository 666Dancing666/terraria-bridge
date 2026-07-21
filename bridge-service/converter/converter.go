package converter

import (
"terraria-bridge/mapping"
"terraria-bridge/protocol"
)

func ConvertWorldSnapshot(original protocol.Message) protocol.Message {
payload := original.Payload

tiles, ok := payload["tiles"].([]interface{})
if !ok {
return original
}

converted := make([]map[string]interface{}, 0, len(tiles))
for _, t := range tiles {
tile, ok := t.(map[string]interface{})
if !ok {
continue
}

newTile := make(map[string]interface{})
newTile["x"] = tile["x"]
newTile["y"] = tile["y"]

if tileType, ok := tile["tile_type"].(float64); ok {
newTile["mc_block"] = mapping.TRTileToMC(int(tileType))
newTile["tr_tile"] = int(tileType)
}

if wallType, ok := tile["wall_type"].(float64); ok {
newTile["mc_wall"] = mapping.TRWallToMC(int(wallType))
newTile["tr_wall"] = int(wallType)
}

if liquidType, ok := tile["liquid_type"].(float64); ok {
newTile["liquid_type"] = int(liquidType)
newTile["liquid_amount"] = tile["liquid_amount"]
}

converted = append(converted, newTile)
}

return protocol.Message{
Type: "world_snapshot_converted",
Payload: map[string]interface{}{
"center_x":    payload["center_x"],
"center_y":    payload["center_y"],
"radius_x":    payload["radius_x"],
"radius_y":    payload["radius_y"],
"world_width":  payload["world_width"],
"world_height": payload["world_height"],
"tiles":       converted,
},
}
}

func ConvertEntityUpdate(original protocol.Message) protocol.Message {
payload := original.Payload

entityType, _ := payload["entity_type"].(string)
converted := make(map[string]interface{})

for k, v := range payload {
converted[k] = v
}

switch entityType {
case "npc":
if netID, ok := payload["net_id"].(float64); ok {
converted["mc_entity"] = mapping.TREntityToMC(int(netID))
converted["tr_entity"] = int(netID)
}
case "item":
if netID, ok := payload["net_id"].(float64); ok {
converted["mc_item"] = mapping.TRItemToMC(int(netID))
converted["tr_item"] = int(netID)
}
case "player":
converted["mc_entity"] = "minecraft:player"
case "projectile":
converted["mc_entity"] = "minecraft:arrow"
}

return protocol.Message{
Type:    "entity_update_converted",
Payload: converted,
}
}

func ConvertTileUpdate(original protocol.Message) protocol.Message {
payload := original.Payload
converted := make(map[string]interface{})

for k, v := range payload {
converted[k] = v
}

if tileType, ok := payload["tile_type"].(float64); ok {
converted["mc_block"] = mapping.TRTileToMC(int(tileType))
}
if wallType, ok := payload["wall_type"].(float64); ok {
converted["mc_wall"] = mapping.TRWallToMC(int(wallType))
}

return protocol.Message{
Type:    "tile_update_converted",
Payload: converted,
}
}

func ConvertPlayerMove(original protocol.Message) protocol.Message {
payload := original.Payload
converted := make(map[string]interface{})

if x, ok := payload["x"].(float64); ok {
converted["x"] = x / 16.0
}
if y, ok := payload["y"].(float64); ok {
converted["y"] = y / 16.0
}
if id, ok := payload["id"]; ok {
converted["id"] = id
}

converted["z"] = 0

return protocol.Message{
Type:    "mc_player_move",
Payload: converted,
}
}

func ConvertMCToTR(original protocol.Message) protocol.Message {
payload := original.Payload
converted := make(map[string]interface{})

for k, v := range payload {
converted[k] = v
}

switch original.Type {
case "player_move":
if x, ok := payload["x"].(float64); ok {
converted["x"] = x * 16.0
}
if y, ok := payload["y"].(float64); ok {
converted["y"] = y * 16.0
}
case "tile_break":
if mcBlock, ok := payload["mc_block"].(string); ok {
converted["tile_type"] = mapping.MCTileToTR(mcBlock)
}
case "tile_place":
if mcBlock, ok := payload["mc_block"].(string); ok {
converted["tile_type"] = mapping.MCTileToTR(mcBlock)
}
}

return protocol.Message{
Type:    original.Type,
Payload: converted,
}
}

func ConvertChatTRToMC(original protocol.Message) protocol.Message {
payload := original.Payload
playerName, _ := payload["player_name"].(string)
msg, _ := payload["msg"].(string)
world, _ := payload["world"].(string)

mcMsg := ""
if world != "" {
mcMsg = "[TR/" + world + "] "
}
if playerName != "" {
mcMsg += "<" + playerName + "> "
}
mcMsg += msg

return protocol.Message{
Type: "chat_message",
Payload: map[string]interface{}{
"msg":         mcMsg,
"source":      "terraria",
"player_name": playerName,
"raw":         msg,
},
}
}

func ConvertChatMCToTR(original protocol.Message) protocol.Message {
payload := original.Payload
playerID, _ := payload["_player_id"].(string)
msg, _ := payload["msg"].(string)

trMsg := "[MC] "
if playerID != "" {
short := playerID
if len(short) > 8 {
short = short[:8]
}
trMsg += "<" + short + "> "
}
trMsg += msg

return protocol.Message{
Type: "chat_message",
Payload: map[string]interface{}{
"msg":    trMsg,
"source": "minecraft",
"player": playerID,
"raw":    msg,
},
}
}
