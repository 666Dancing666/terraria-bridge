using Terraria;

namespace TerrariaBridge.Utils
{
    public static class TileHelper
    {
        public static void BreakTile(int x, int y, int playerId)
        {
            if (!WorldGen.InWorld(x, y)) return;

            // 先尝试破坏瓦片
            if (Main.tile[x, y] != null && Main.tile[x, y].HasTile)
            {
                WorldGen.KillTile(x, y, false, false, false);
                if (Main.netMode == 2)
                {
                    NetMessage.SendTileSquare(-1, x, y, 1);
                }
            }
            // 再尝试破坏墙壁
            else if (Main.tile[x, y] != null && Main.tile[x, y].WallType > 0)
            {
                WorldGen.KillWall(x, y, false);
                if (Main.netMode == 2)
                {
                    NetMessage.SendTileSquare(-1, x, y, 1);
                }
            }
        }

        public static void PlaceTile(int x, int y, int tileType, int playerId)
        {
            if (!WorldGen.InWorld(x, y)) return;

            // 如果已经有瓦片，不放置
            if (Main.tile[x, y] != null && Main.tile[x, y].HasTile) return;

            WorldGen.PlaceTile(x, y, tileType, false, true, playerId);
            if (Main.netMode == 2)
            {
                NetMessage.SendTileSquare(-1, x, y, 1);
            }
        }

        public static void PlaceWall(int x, int y, int wallType, int playerId)
        {
            if (!WorldGen.InWorld(x, y)) return;

            // 如果已经有墙，不放置
            if (Main.tile[x, y] != null && Main.tile[x, y].WallType > 0) return;

            WorldGen.PlaceWall(x, y, wallType, false);
            if (Main.netMode == 2)
            {
                NetMessage.SendTileSquare(-1, x, y, 1);
            }
        }

        public static bool IsTileBlocked(int x, int y)
        {
            if (!WorldGen.InWorld(x, y)) return true;

            var tile = Main.tile[x, y];
            if (tile == null) return true;

            // 被实心瓦片挡住
            if (tile.HasTile && Main.tileSolid[tile.TileType]) return true;

            return false;
        }
    }
}
