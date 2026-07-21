using System.Collections.Generic;
using System.Linq;
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

            int startX = (centerX - radiusX).Clamp(0, Main.maxTilesX - 1);
            int endX = (centerX + radiusX).Clamp(0, Main.maxTilesX - 1);
            int startY = (centerY - radiusY).Clamp(0, Main.maxTilesY - 1);
            int endY = (centerY + radiusY).Clamp(0, Main.maxTilesY - 1);

            for (int x = startX; x <= endX; x++)
            {
                for (int y = startY; y <= endY; y++)
                {
                    var tile = Main.tile[x, y];
                    if (tile == null) continue;

                    bool hasData = tile.HasTile || tile.WallType > 0 || tile.LiquidAmount > 0;

                    if (!hasData) continue;

                    var tileData = new Dictionary<string, object>
                    {
                        { "x", x },
                        { "y", y }
                    };

                    if (tile.HasTile)
                    {
                        tileData["tile_type"] = tile.TileType;
                        tileData["tile_frame_x"] = tile.TileFrameX;
                        tileData["tile_frame_y"] = tile.TileFrameY;
                        tileData["tile_color"] = tile.TileColor;
                        tileData["half_brick"] = tile.IsHalfBlock;
                        tileData["slope"] = (byte)tile.Slope;
                    }

                    if (tile.WallType > 0)
                    {
                        tileData["wall_type"] = tile.WallType;
                        tileData["wall_frame_x"] = tile.WallFrameX;
                        tileData["wall_frame_y"] = tile.WallFrameY;
                        tileData["wall_color"] = tile.WallColor;
                    }

                    if (tile.LiquidAmount > 0)
                    {
                        tileData["liquid_type"] = tile.LiquidType;
                        tileData["liquid_amount"] = tile.LiquidAmount;
                    }

                    tiles.Add(tileData);
                }
            }

            var snapshot = new BridgeMessage
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

            return snapshot;
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
                    { "tile_type", tile.TileType },
                    { "tile_frame_x", tile.TileFrameX },
                    { "tile_frame_y", tile.TileFrameY },
                    { "tile_color", tile.TileColor },
                    { "half_brick", tile.IsHalfBlock },
                    { "slope", (byte)tile.Slope },
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
                    { "wall_type", tile.WallType },
                    { "wall_frame_x", tile.WallFrameX },
                    { "wall_frame_y", tile.WallFrameY },
                    { "wall_color", tile.WallColor }
                }
            };
        }
    }
}

public static class IntExtensions
{
    public static int Clamp(this int value, int min, int max)
    {
        if (value < min) return min;
        if (value > max) return max;
        return value;
    }
}
