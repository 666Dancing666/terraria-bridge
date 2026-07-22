using Terraria;

namespace TerrariaBridge.Utils
{
    public static class TileHelper
    {
        public static void BreakTile(int x, int y, int playerId)
        {
            if (!WorldGen.InWorld(x, y)) return;
            Tile tile = Main.tile[x, y];
            if (tile == null) return;

            WorldGen.KillTile(x, y, false, false, false);
            if (Main.netMode == 2)
            {
                NetMessage.SendTileSquare(-1, x, y, 1);
            }
        }

        public static void PlaceTile(int x, int y, int tileType, int playerId)
        {
            if (!WorldGen.InWorld(x, y)) return;
            WorldGen.PlaceTile(x, y, tileType, false, true, playerId);
            if (Main.netMode == 2)
            {
                NetMessage.SendTileSquare(-1, x, y, 1);
            }
        }

        public static void PlaceWall(int x, int y, int wallType, int playerId)
        {
            if (!WorldGen.InWorld(x, y)) return;
            WorldGen.PlaceWall(x, y, wallType, false);
            if (Main.netMode == 2)
            {
                NetMessage.SendTileSquare(-1, x, y, 1);
            }
        }

        public static bool IsTileBlocked(int x, int y)
        {
            if (!WorldGen.InWorld(x, y)) return true;
            Tile tile = Main.tile[x, y];
            if (tile == null) return true;
            if (tile.active() && Main.tileSolid[tile.type]) return true;
            return false;
        }
    }
}
