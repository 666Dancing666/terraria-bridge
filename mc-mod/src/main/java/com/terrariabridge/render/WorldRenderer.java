package com.terrariabridge.render;

import net.minecraft.client.Minecraft;
import net.minecraft.world.level.Level;
import net.minecraftforge.event.TickEvent;
import net.minecraftforge.eventbus.api.SubscribeEvent;

public class WorldRenderer
{
    private final LayerManager layerManager;
    private int lastCenterX = 0;
    private int lastCenterY = 0;
    private static final int RENDER_RADIUS = 100;

    public WorldRenderer(LayerManager layerManager)
    {
        this.layerManager = layerManager;
    }

    @SubscribeEvent
    public void onClientTick(TickEvent.ClientTickEvent event)
    {
        if (event.phase != TickEvent.Phase.END) return;

        Minecraft mc = Minecraft.getInstance();
        if (mc.player == null || mc.level == null) return;

        int trX = (int) mc.player.getX();
        int trY = (int) mc.player.getY();

        if (Math.abs(trX - lastCenterX) > 20 ||
            Math.abs(trY - lastCenterY) > 20)
        {
            lastCenterX = trX;
            lastCenterY = trY;

            int minX = trX - RENDER_RADIUS;
            int minY = trY - RENDER_RADIUS;
            int maxX = trX + RENDER_RADIUS;
            int maxY = trY + RENDER_RADIUS;

            layerManager.applyToWorld(mc.level, minX, minY, maxX, maxY);
        }
    }
}
