using Terraria;
using Terraria.ID;

namespace TerrariaBridge.Utils
{
    public static class TileHelper
    {
        public static void BreakTile(int x, int y, int playerId)
        {
            if (!WorldGen.InWorld(x, y)) return;
            var tile = Main.tile[x, y];
            if (tile == null) return;

            if (tile.HasTile)
            {
                WorldGen.KillTile(x, y, false, false, false);
                if (Main.netMode == NetmodeID.Server)
                {
                    NetMessage.SendTileSquare(-1, x, y, 1);
                }
            }
            else if (tile.WallType > 0)
            {
                WorldGen.KillWall(x, y, false);
                if (Main.netMode == NetmodeID.Server)
                {
                    NetMessage.SendTileSquare(-1, x, y, 1);
                }
            }
        }

        public static void PlaceTile(int x, int y, int tileType, int playerId)
        {
            if (!WorldGen.InWorld(x, y)) return;
            var tile = Main.tile[x, y];
            if (tile == null) return;
            if (tile.HasTile) return;

            WorldGen.PlaceTile(x, y, tileType, false, true, playerId);
            if (Main.netMode == NetmodeID.Server)
            {
                NetMessage.SendTileSquare(-1, x, y, 1);
            }
        }

        public static void PlaceWall(int x, int y, int wallType, int playerId)
        {
            if (!WorldGen.InWorld(x, y)) return;
            var tile = Main.tile[x, y];
            if (tile == null) return;
            if (tile.WallType > 0) return;

            WorldGen.PlaceWall(x, y, wallType, false);
            if (Main.netMode == NetmodeID.Server)
            {
                NetMessage.SendTileSquare(-1, x, y, 1);
            }
        }

        public static bool IsTileBlocked(int x, int y)
        {
            if (!WorldGen.InWorld(x, y)) return true;
            var tile = Main.tile[x, y];
            if (tile == null) return true;
            if (tile.HasTile && Main.tileSolid[tile.TileType]) return true;
            return false;
        }
    }
}
