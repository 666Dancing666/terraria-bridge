using System.Collections.Generic;
using Terraria;
using TerrariaBridge.Network;

namespace TerrariaBridge.Handlers
{
    public class WorldSyncHandler
    {
        private BridgeClient _client;

        public WorldSyncHandler(BridgeClient client)
        {
            _client = client;
        }

        public BridgeMessage CreateFullSnapshot(int centerX, int centerY, int radiusX, int radiusY)
        {
            var tiles = new List<Dictionary<string, object>>();

            int startX = Clamp(centerX - radiusX, 0, Main.maxTilesX - 1);
            int endX = Clamp(centerX + radiusX, 0, Main.maxTilesX - 1);
            int startY = Clamp(centerY - radiusY, 0, Main.maxTilesY - 1);
            int endY = Clamp(centerY + radiusY, 0, Main.maxTilesY - 1);

            for (int x = startX; x <= endX; x++)
            {
                for (int y = startY; y <= endY; y++)
                {
                    var tile = Main.tile[x, y];
                    if (tile == null) continue;

                    var tileData = new Dictionary<string, object>
                    {
                        { "x", x },
                        { "y", y }
                    };

                    if (tile.HasTile)
                    {
                        tileData["tile_type"] = (int)tile.TileType;
                        tileData["has_tile"] = true;
                    }

                    if (tile.WallType > 0)
                    {
                        tileData["wall_type"] = (int)tile.WallType;
                    }

                    if (tile.LiquidAmount > 0)
                    {
                        tileData["liquid_type"] = (int)tile.LiquidType;
                        tileData["liquid_amount"] = (int)tile.LiquidAmount;
                    }

                    if (tileData.Count > 2)
                    {
                        tiles.Add(tileData);
                    }
                }
            }

            return new BridgeMessage
            {
                Type = MessageType.WorldSnapshot,
                Payload = new Dictionary<string, object>
                {
                    { "center_x", centerX },
                    { "center_y", centerY },
                    { "radius_x", radiusX },
                    { "radius_y", radiusY },
                    { "world_width", Main.maxTilesX },
                    { "world_height", Main.maxTilesY },
                    { "tiles", tiles }
                }
            };
        }

        public BridgeMessage CreateTileUpdate(int x, int y)
        {
            var tile = Main.tile[x, y];
            if (tile == null) return null;

            return new BridgeMessage
            {
                Type = MessageType.TileUpdate,
                Payload = new Dictionary<string, object>
                {
                    { "x", x },
                    { "y", y },
                    { "tile_type", (int)tile.TileType },
                    { "has_tile", tile.HasTile }
                }
            };
        }

        public BridgeMessage CreateWallUpdate(int x, int y)
        {
            var tile = Main.tile[x, y];
            if (tile == null) return null;

            return new BridgeMessage
            {
                Type = MessageType.WallUpdate,
                Payload = new Dictionary<string, object>
                {
                    { "x", x },
                    { "y", y },
                    { "wall_type", (int)tile.WallType }
                }
            };
        }

        private static int Clamp(int value, int min, int max)
        {
            if (value < min) return min;
            if (value > max) return max;
            return value;
        }
    }
}
