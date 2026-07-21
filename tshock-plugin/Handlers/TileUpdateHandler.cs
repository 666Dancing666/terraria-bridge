using System;
using System.Collections.Generic;
using Terraria;
using TerrariaApi.Server;
using TShockAPI;
using TerrariaBridge.Network;
using TerrariaBridge.Handlers;

namespace TerrariaBridge.Handlers
{
    public class TileUpdateHandler
    {
        private BridgeClient _client;
        private WorldSyncHandler _worldSync;
        private bool[,] _trackedTiles;
        private int _lastTrackX = -1;
        private int _lastTrackY = -1;
        private int _trackRadius = 120;

        public TileUpdateHandler(BridgeClient client, WorldSyncHandler worldSync)
        {
            _client = client;
            _worldSync = worldSync;
            _trackedTiles = new bool[Main.maxTilesX, Main.maxTilesY];
        }

        public void Register()
        {
            ServerApi.Hooks.GameUpdate.Register(TerrariaBridgePlugin.Instance, OnGameUpdate);
        }

        private void OnGameUpdate(GameUpdateEventArgs args)
        {
            if (!_client.IsConnected) return;

            int centerX = 0;
            int centerY = 0;
            bool found = false;

            for (int i = 0; i < Main.player.Length; i++)
            {
                if (Main.player[i] != null && Main.player[i].active)
                {
                    centerX = (int)(Main.player[i].position.X / 16);
                    centerY = (int)(Main.player[i].position.Y / 16);
                    found = true;
                    break;
                }
            }

            if (!found) return;

            int minX = Math.Max(0, centerX - _trackRadius);
            int maxX = Math.Min(Main.maxTilesX - 1, centerX + _trackRadius);
            int minY = Math.Max(0, centerY - _trackRadius);
            int maxY = Math.Min(Main.maxTilesY - 1, centerY + _trackRadius);

            for (int x = minX; x <= maxX; x++)
            {
                for (int y = minY; y <= maxY; y++)
                {
                    var tile = Main.tile[x, y];
                    if (tile == null) continue;

                    bool hasTile = tile.HasTile;
                    bool wasTracked = _trackedTiles[x, y];

                    if (hasTile != wasTracked)
                    {
                        _trackedTiles[x, y] = hasTile;

                        if (hasTile && tile.TileType > 0)
                        {
                            var msg = _worldSync.CreateTileUpdate(x, y);
                            if (msg != null)
                            {
                                _client.SendAsync(msg).Wait();
                            }
                        }
                        else
                        {
                            var msg = new BridgeMessage
                            {
                                Type = MessageType.TileUpdate,
                                Payload = new Dictionary<string, object>
                                {
                                    { "x", x },
                                    { "y", y },
                                    { "tile_type", 0 },
                                    { "has_tile", false }
                                }
                            };
                            _client.SendAsync(msg).Wait();
                        }
                    }

                    if (hasTile && wasTracked)
                    {
                        int oldType = 0;
                        if (_trackedTiles[x, y] && tile.TileType != oldType)
                        {
                            var msg = _worldSync.CreateTileUpdate(x, y);
                            if (msg != null)
                            {
                                _client.SendAsync(msg).Wait();
                            }
                        }
                    }

                    bool hasWall = tile.WallType > 0;
                    if (hasWall != wasTracked)
                    {
                        var wallMsg = _worldSync.CreateWallUpdate(x, y);
                        if (wallMsg != null)
                        {
                            _client.SendAsync(wallMsg).Wait();
                        }
                    }
                }
            }
        }
    }
}
