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
                case MessageType.Interact:
                    HandleInteract(msg.Payload);
                    break;
                case MessageType.PlayerAction:
                    HandlePlayerAction(msg.Payload);
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
            int playerId = payload.ContainsKey("player_id") 
                ? Convert.ToInt32(payload["player_id"]) : -1;

            TileHelper.BreakTile(x, y, playerId);
        }

        private void HandleTilePlace(Dictionary<string, object> payload)
        {
            int x = Convert.ToInt32(payload["x"]);
            int y = Convert.ToInt32(payload["y"]);
            int tileType = Convert.ToInt32(payload["tile_type"]);
            int playerId = payload.ContainsKey("player_id") 
                ? Convert.ToInt32(payload["player_id"]) : -1;

            TileHelper.PlaceTile(x, y, tileType, playerId);
        }

        private void HandleInteract(Dictionary<string, object> payload)
        {
            int x = Convert.ToInt32(payload["x"]);
            int y = Convert.ToInt32(payload["y"]);
            int playerId = Convert.ToInt32(payload["player_id"]);

            // 打开箱子或交互
            if (playerId >= 0 && playerId < Main.player.Length)
            {
                var player = Main.player[playerId];
                if (player != null && player.active)
                {
                    // 让玩家与被点击的物体交互
                    player.tileInteractAttempted = true;
                    player.tileInteractionX = x;
                    player.tileInteractionY = y;
                }
            }
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
                    switch (action)
                    {
                        case "jump":
                            if (player.velocity.Y == 0)
                            {
                                player.velocity.Y = -player.jumpSpeed;
                            }
                            break;
                        case "use_item":
                            player.controlUseItem = true;
                            break;
                        case "grapple":
                            player.controlHook = true;
                            break;
                    }
                }
            }
        }
    }
}
