using System;
using System.Collections.Generic;
using Terraria;
using TerrariaBridge.Network;
using TerrariaBridge.Utils;

namespace TerrariaBridge.Handlers
{
    public class CommandHandler
    {
        private BridgeClient _client;

        public CommandHandler(BridgeClient client)
        {
            _client = client;
        }

        public void HandleMessage(BridgeMessage msg)
        {
            switch (msg.Type)
            {
                case MessageType.PlayerMove:
                    HandlePlayerMove(msg.Payload);
                    break;
                case MessageType.TileBreak:
                    HandleTileBreak(msg.Payload);
                    break;
                case MessageType.TilePlace:
                    HandleTilePlace(msg.Payload);
                    break;
                case MessageType.PlayerAction:
                    HandlePlayerAction(msg.Payload);
                    break;
                case "craft_request":
                    HandleCraftRequest(msg.Payload);
                    break;
            }
        }

        private void HandlePlayerMove(Dictionary<string, object> payload)
        {
            int playerId = Convert.ToInt32(payload["id"]);
            float x = Convert.ToSingle(payload["x"]);
            float y = Convert.ToSingle(payload["y"]);

            if (playerId >= 0 && playerId < Main.player.Length)
            {
                var player = Main.player[playerId];
                if (player != null && player.active)
                {
                    player.position.X = x;
                    player.position.Y = y;
                    player.velocity.X = 0;
                    player.velocity.Y = 0;
                }
            }
        }

        private void HandleTileBreak(Dictionary<string, object> payload)
        {
            int x = Convert.ToInt32(payload["x"]);
            int y = Convert.ToInt32(payload["y"]);
            TileHelper.BreakTile(x, y);
        }

        private void HandleTilePlace(Dictionary<string, object> payload)
        {
            int x = Convert.ToInt32(payload["x"]);
            int y = Convert.ToInt32(payload["y"]);
            int tileType = Convert.ToInt32(payload["tile_type"]);
            TileHelper.PlaceTile(x, y, tileType);
        }

        private void HandlePlayerAction(Dictionary<string, object> payload)
        {
            int playerId = Convert.ToInt32(payload["player_id"]);
            string action = payload["action"].ToString();

            if (playerId >= 0 && playerId < Main.player.Length)
            {
                var player = Main.player[playerId];
                if (player != null && player.active)
                {
                    if (action == "jump")
                    {
                        player.controlJump = true;
                        player.releaseJump = false;
                    }
                    else if (action == "use_item")
                    {
                        player.controlUseItem = true;
                    }
                    else if (action == "grapple")
                    {
                        float targetX = player.position.X + player.direction * 160;
                        float targetY = player.position.Y - 80;
                        player.position.X = targetX;
                        player.position.Y = targetY;
                        player.velocity.X = 0;
                        player.velocity.Y = 0;
                    }
                }
            }
        }

        private void HandleCraftRequest(Dictionary<string, object> payload)
        {
            int playerId = 0;
            int resultId = Convert.ToInt32(payload["result_id"]);
            int resultCount = Convert.ToInt32(payload["result_count"]);

            if (playerId >= 0 && playerId < Main.player.Length)
            {
                var player = Main.player[playerId];
                if (player == null || !player.active) return;

                var ingredients = payload["ingredients"] as Dictionary<string, object>;
                if (ingredients == null) return;

                var needed = new Dictionary<int, int>();
                foreach (var kvp in ingredients)
                {
                    int itemId = int.Parse(kvp.Key);
                    int count = Convert.ToInt32(kvp.Value);
                    needed[itemId] = count;
                }

                bool hasAll = true;
                foreach (var kvp in needed)
                {
                    int count = 0;
                    for (int i = 0; i < player.inventory.Length; i++)
                    {
                        if (player.inventory[i].type == kvp.Key)
                            count += player.inventory[i].stack;
                    }
                    if (count < kvp.Value) { hasAll = false; break; }
                }

                if (hasAll)
                {
                    foreach (var kvp in needed)
                    {
                        int remaining = kvp.Value;
                        for (int i = 0; i < player.inventory.Length && remaining > 0; i++)
                        {
                            if (player.inventory[i].type == kvp.Key)
                            {
                                int take = Math.Min(remaining, player.inventory[i].stack);
                                player.inventory[i].stack -= take;
                                remaining -= take;
                                if (player.inventory[i].stack <= 0)
                                    player.inventory[i].TurnToAir();
                            }
                        }
                    }
                    player.QuickSpawnItem(null, resultId, resultCount);
                }
            }
        }
    }
}
