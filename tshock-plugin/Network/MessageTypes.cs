using System.Collections.Generic;
using Newtonsoft.Json;

namespace TerrariaBridge.Network
{
    public class BridgeMessage
    {
        [JsonProperty("type")]
        public string Type { get; set; }

        [JsonProperty("payload")]
        public Dictionary<string, object> Payload { get; set; }
            = new Dictionary<string, object>();
    }

    public static class MessageType
    {
        // TShock → MC
        public const string WorldSnapshot = "world_snapshot";
        public const string TileUpdate = "tile_update";
        public const string WallUpdate = "wall_update";
        public const string LiquidUpdate = "liquid_update";
        public const string EntityUpdate = "entity_update";
        public const string EntityRemove = "entity_remove";
        public const string PlayerJoin = "player_join";
        public const string PlayerLeave = "player_leave";
        public const string ChatMessage = "chat_message";

        // MC → TShock
        public const string PlayerMove = "player_move";
        public const string PlayerAction = "player_action";
        public const string TileBreak = "tile_break";
        public const string TilePlace = "tile_place";
        public const string Interact = "interact";
    }
}
