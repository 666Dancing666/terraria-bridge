using Terraria;

namespace TerrariaBridge.Utils
{
    public static class TileHelper
    {
        public static void BreakTile(int x, int y)
        {
            if (!WorldGen.InWorld(x, y)) return;
            WorldGen.KillTile(x, y);
            if (Main.netMode == 2)
                NetMessage.SendTileSquare(-1, x, y, 1);
        }

        public static void PlaceTile(int x, int y, int tileType)
        {
            if (!WorldGen.InWorld(x, y)) return;
            WorldGen.PlaceTile(x, y, tileType, true, true);
            if (Main.netMode == 2)
                NetMessage.SendTileSquare(-1, x, y, 1);
        }

        public static void PlaceWall(int x, int y, int wallType)
        {
            if (!WorldGen.InWorld(x, y)) return;
            WorldGen.PlaceWall(x, y, wallType, true);
            if (Main.netMode == 2)
                NetMessage.SendTileSquare(-1, x, y, 1);
        }
    }
}
