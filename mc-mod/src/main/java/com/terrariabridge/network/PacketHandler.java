package com.terrariabridge.network;

import com.terrariabridge.network.MessageTypes.BridgeMessage;
import com.terrariabridge.render.LayerManager;
import net.minecraft.client.Minecraft;
import net.minecraft.network.chat.Component;
import java.util.*;

public class PacketHandler
{
    private LayerManager layerManager;
    private long lastTimeUpdate = 0;

    public void setLayerManager(LayerManager layerManager)
    {
        this.layerManager = layerManager;
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
            default:
                break;
        }
    }

    @SuppressWarnings("unchecked")
    private void handleWorldSnapshot(Map<String, Object> payload)
    {
        if (layerManager == null) return;
        int centerX = ((Number) payload.get("center_x")).intValue();
        int centerY = ((Number) payload.get("center_y")).intValue();
        System.out.println("TerrariaBridge: Received world snapshot, center(" +
            centerX + "," + centerY + ")");
        List<Map<String, Object>> tiles = (List<Map<String, Object>>) payload.get("tiles");
        if (tiles != null) {
            for (Map<String, Object> tile : tiles) {
                layerManager.updateTile(tile);
            }
        }
    }

    private void handleTileUpdate(Map<String, Object> payload) {
        if (layerManager != null) layerManager.updateTile(payload);
    }

    private void handleWallUpdate(Map<String, Object> payload) {
        if (layerManager != null) layerManager.updateTile(payload);
    }

    private void handleLiquidUpdate(Map<String, Object> payload) {
        if (layerManager != null) layerManager.updateTile(payload);
    }

    private void handleEntityUpdate(Map<String, Object> payload) {
        if (layerManager != null) layerManager.updateEntity(payload);
    }

    private void handleEntityRemove(Map<String, Object> payload) {
        if (layerManager != null) layerManager.removeEntity(payload);
    }

    private void handleTimeSync(Map<String, Object> payload) {
        long now = System.currentTimeMillis();
        if (now - lastTimeUpdate < 5000) return;
        lastTimeUpdate = now;

        Minecraft mc = Minecraft.getInstance();
        if (mc.level != null) {
            long mcTime = ((Number) payload.get("mc_time")).longValue();
            mc.level.setDayTime(mcTime);
        }
    }
}
