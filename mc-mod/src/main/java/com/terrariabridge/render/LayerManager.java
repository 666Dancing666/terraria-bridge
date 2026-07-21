package com.terrariabridge.render;

import net.minecraft.core.BlockPos;
import net.minecraft.world.level.Level;
import net.minecraft.world.level.block.Blocks;
import java.util.*;
import java.util.concurrent.ConcurrentHashMap;

public class LayerManager
{
    public static final int LAYER_FOREGROUND = 0;
    public static final int LAYER_FURNITURE = 1;
    public static final int LAYER_WALL = 2;

    private final Map<Long, LayerData> tileData = new ConcurrentHashMap<>();
    private final Map<Integer, EntityData> entities = new ConcurrentHashMap<>();

    private static class LayerData
    {
        int tileType;
        int wallType;
        int furnitureType;
        int liquidType;
        int liquidAmount;
        boolean hasTile;
    }

    public static class EntityData
    {
        public String entityType;
        public int id;
        public String name;
        public double x, y;
        public double velocityX, velocityY;
        public int health, maxHealth;
        public int direction;
        public int netId;
    }

    public void updateTile(Map<String, Object> data)
    {
        int x = ((Number) data.get("x")).intValue();
        int y = ((Number) data.get("y")).intValue();
        long key = posKey(x, y);

        LayerData layer = tileData.computeIfAbsent(key, k -> new LayerData());

        if (data.containsKey("tile_type"))
        {
            int type = ((Number) data.get("tile_type")).intValue();
            boolean has = !data.containsKey("has_tile") || (Boolean) data.get("has_tile");
            if (has && type > 0) { layer.tileType = type; layer.hasTile = true; }
            else { layer.tileType = 0; layer.hasTile = false; }
        }

        if (data.containsKey("wall_type"))
        {
            layer.wallType = ((Number) data.get("wall_type")).intValue();
        }

        if (data.containsKey("liquid_type"))
        {
            layer.liquidType = ((Number) data.get("liquid_type")).intValue();
            layer.liquidAmount = ((Number) data.get("liquid_amount")).intValue();
        }
    }

    public void updateEntity(Map<String, Object> data)
    {
        int id = ((Number) data.get("id")).intValue();
        EntityData entity = entities.computeIfAbsent(id, k -> new EntityData());
        entity.entityType = (String) data.get("entity_type");
        entity.id = id;
        entity.name = (String) data.get("name");
        entity.x = ((Number) data.get("x")).doubleValue();
        entity.y = ((Number) data.get("y")).doubleValue();
        if (data.containsKey("direction")) entity.direction = ((Number) data.get("direction")).intValue();
        if (data.containsKey("health")) entity.health = ((Number) data.get("health")).intValue();
        if (data.containsKey("max_health")) entity.maxHealth = ((Number) data.get("max_health")).intValue();
        if (data.containsKey("net_id")) entity.netId = ((Number) data.get("net_id")).intValue();
    }

    public void removeEntity(Map<String, Object> data)
    {
        int id = ((Number) data.get("id")).intValue();
        entities.remove(id);
    }

    public int getTileType(int x, int y)
    {
        LayerData layer = tileData.get(posKey(x, y));
        return (layer != null && layer.hasTile) ? layer.tileType : 0;
    }

    public int getWallType(int x, int y)
    {
        LayerData layer = tileData.get(posKey(x, y));
        return layer != null ? layer.wallType : 0;
    }

    public Collection<EntityData> getAllEntities() { return entities.values(); }
    public EntityData getEntity(int id) { return entities.get(id); }

    private static long posKey(int x, int y) { return ((long) x << 32) | (y & 0xFFFFFFFFL); }

    public void applyToWorld(Level level, int minX, int minY, int maxX, int maxY)
    {
        for (int x = minX; x <= maxX; x++)
        {
            for (int y = minY; y <= maxY; y++)
            {
                LayerData layer = tileData.get(posKey(x, y));
                int tile = (layer != null && layer.hasTile) ? layer.tileType : 0;
                int wall = (layer != null) ? layer.wallType : 0;

                level.setBlock(new BlockPos(x, y, LAYER_FOREGROUND),
                    tile > 0 ? Blocks.STONE.defaultBlockState() : Blocks.AIR.defaultBlockState(), 3);
                level.setBlock(new BlockPos(x, y, LAYER_FURNITURE), Blocks.AIR.defaultBlockState(), 3);
                level.setBlock(new BlockPos(x, y, LAYER_WALL),
                    wall > 0 ? Blocks.GRAY_WOOL.defaultBlockState() : Blocks.AIR.defaultBlockState(), 3);
            }
        }
    }
}
