package com.terrariabridge.network;

import com.terrariabridge.network.MessageTypes.BridgeMessage;
import com.terrariabridge.render.LayerManager;
import com.terrariabridge.render.WorldRenderer;
import net.minecraft.client.Minecraft;
import java.util.*;

public class PacketHandler
{
    private LayerManager layerManager;
    private WorldRenderer worldRenderer;
    private long lastTimeUpdate = 0;
    private long lastWeatherUpdate = 0;

    public void setLayerManager(LayerManager layerManager)
    {
        this.layerManager = layerManager;
    }

    public void setWorldRenderer(WorldRenderer worldRenderer)
    {
        this.worldRenderer = worldRenderer;
    }

    public void handleMessage(String json)
    {
        BridgeMessage msg = MessageTypes.parse(json);
        if (msg == null) return;

        switch (msg.type)
        {
            case MessageTypes.WORLD_SNAPSHOT:
                handleWorldSnapshot(msg.payload);
                break;
            case MessageTypes.TILE_UPDATE:
                handleTileUpdate(msg.payload);
                break;
            case MessageTypes.WALL_UPDATE:
                handleWallUpdate(msg.payload);
                break;
            case MessageTypes.LIQUID_UPDATE:
                handleLiquidUpdate(msg.payload);
                break;
            case MessageTypes.ENTITY_UPDATE:
                handleEntityUpdate(msg.payload);
                break;
            case MessageTypes.ENTITY_REMOVE:
                handleEntityRemove(msg.payload);
                break;
            case "time_sync":
                handleTimeSync(msg.payload);
                break;
            case "weather_sync":
                handleWeatherSync(msg.payload);
                break;
            default:
                break;
        }
    }

    @SuppressWarnings("unchecked")
    private void handleWorldSnapshot(Map<String, Object> payload)
    {
        if (layerManager == null) return;
        List<Map<String, Object>> tiles = (List<Map<String, Object>>) payload.get("tiles");
        if (tiles != null) {
            for (Map<String, Object> tile : tiles) {
                layerManager.updateTile(tile);
            }
        }
    }

    private void handleTileUpdate(Map<String, Object> payload)
    {
        if (layerManager == null) return;
        layerManager.updateTile(payload);

        if (worldRenderer != null)
        {
            int x = ((Number) payload.get("x")).intValue();
            int y = ((Number) payload.get("y")).intValue();
            String mcBlock = (String) payload.get("mc_block");
            boolean hasTile = !payload.containsKey("has_tile") || (Boolean) payload.get("has_tile");

            if (hasTile && mcBlock != null)
            {
                worldRenderer.scheduleBlockUpdate(x, y, 0, mcBlock);
            }
            else
            {
                worldRenderer.scheduleBlockUpdate(x, y, 0, null);
            }
        }
    }

    private void handleWallUpdate(Map<String, Object> payload)
    {
        if (layerManager == null) return;
        layerManager.updateTile(payload);

        if (worldRenderer != null)
        {
            int x = ((Number) payload.get("x")).intValue();
            int y = ((Number) payload.get("y")).intValue();
            String mcWall = (String) payload.get("mc_wall");
            int wallType = payload.containsKey("wall_type") ? ((Number) payload.get("wall_type")).intValue() : 0;

            if (wallType > 0 && mcWall != null)
            {
                worldRenderer.scheduleBlockUpdate(x, y, 2, mcWall);
            }
            else
            {
                worldRenderer.scheduleBlockUpdate(x, y, 2, null);
            }
        }
    }

    private void handleLiquidUpdate(Map<String, Object> payload)
    {
        if (layerManager != null) layerManager.updateTile(payload);
    }

    private void handleEntityUpdate(Map<String, Object> payload)
    {
        if (layerManager != null) layerManager.updateEntity(payload);
    }

    private void handleEntityRemove(Map<String, Object> payload)
    {
        if (layerManager != null) layerManager.removeEntity(payload);
    }

    private void handleTimeSync(Map<String, Object> payload)
    {
        long now = System.currentTimeMillis();
        if (now - lastTimeUpdate < 5000) return;
        lastTimeUpdate = now;
        Minecraft mc = Minecraft.getInstance();
        if (mc.level != null) {
            long mcTime = ((Number) payload.get("mc_time")).longValue();
            mc.level.setDayTime(mcTime);
        }
    }

    private void handleWeatherSync(Map<String, Object> payload)
    {
        long now = System.currentTimeMillis();
        if (now - lastWeatherUpdate < 10000) return;
        lastWeatherUpdate = now;
        Minecraft mc = Minecraft.getInstance();
        if (mc.level != null) {
            boolean rain = (Boolean) payload.get("mc_rain");
            boolean thunder = (Boolean) payload.get("mc_thunder");
            if (rain) {
                mc.level.setRainLevel(1.0f);
                mc.level.setThunderLevel(thunder ? 1.0f : 0.0f);
            } else {
                mc.level.setRainLevel(0.0f);
                mc.level.setThunderLevel(0.0f);
            }
        }
    }
}
