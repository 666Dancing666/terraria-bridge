package com.terrariabridge.render;

import net.minecraft.client.Minecraft;
import net.minecraft.core.BlockPos;
import net.minecraft.world.level.Level;
import net.minecraft.world.level.block.Block;
import net.minecraft.world.level.block.Blocks;
import net.minecraft.resources.ResourceLocation;
import net.minecraftforge.registries.ForgeRegistries;
import net.minecraftforge.event.TickEvent;
import net.minecraftforge.eventbus.api.SubscribeEvent;
import java.util.*;
import java.util.concurrent.ConcurrentLinkedQueue;

public class WorldRenderer
{
    private final LayerManager layerManager;
    private int lastCenterX = Integer.MAX_VALUE;
    private int lastCenterY = Integer.MAX_VALUE;
    private static final int RENDER_RADIUS = 60;
    private final Queue<BlockUpdate> pendingUpdates = new ConcurrentLinkedQueue<>();
    private boolean initialRender = true;

    private static class BlockUpdate
    {
        int x, y, z;
        String mcName;

        BlockUpdate(int x, int y, int z, String mcName)
        {
            this.x = x;
            this.y = y;
            this.z = z;
            this.mcName = mcName;
        }
    }

    public WorldRenderer(LayerManager layerManager)
    {
        this.layerManager = layerManager;
    }

    public void scheduleBlockUpdate(int x, int y, int z, String mcName)
    {
        pendingUpdates.add(new BlockUpdate(x, y, z, mcName));
    }

    @SubscribeEvent
    public void onClientTick(TickEvent.ClientTickEvent event)
    {
        if (event.phase != TickEvent.Phase.END) return;

        Minecraft mc = Minecraft.getInstance();
        if (mc.player == null || mc.level == null) return;

        Level level = mc.level;

        int trX = (int) mc.player.getX();
        int trY = (int) mc.player.getY();

        if (initialRender || Math.abs(trX - lastCenterX) > 30 || Math.abs(trY - lastCenterY) > 30)
        {
            if (initialRender || Math.abs(trX - lastCenterX) > 15 || Math.abs(trY - lastCenterY) > 15)
            {
                int minX = trX - RENDER_RADIUS;
                int minY = trY - RENDER_RADIUS;
                int maxX = trX + RENDER_RADIUS;
                int maxY = trY + RENDER_RADIUS;

                for (int x = lastCenterX - RENDER_RADIUS - 5; x <= lastCenterX + RENDER_RADIUS + 5; x++)
                {
                    for (int y = lastCenterY - RENDER_RADIUS - 5; y <= lastCenterY + RENDER_RADIUS + 5; y++)
                    {
                        if (x < minX || x > maxX || y < minY || y > maxY)
                        {
                            for (int z = 0; z < 3; z++)
                            {
                                level.setBlock(new BlockPos(x, y, z), Blocks.AIR.defaultBlockState(), 3);
                            }
                        }
                    }
                }

                layerManager.applyToWorld(level, minX, minY, maxX, maxY);
            }

            lastCenterX = trX;
            lastCenterY = trY;
            initialRender = false;
        }

        int processedCount = 0;
        while (!pendingUpdates.isEmpty() && processedCount < 500)
        {
            BlockUpdate update = pendingUpdates.poll();
            processedCount++;

            if (update.x >= lastCenterX - RENDER_RADIUS && update.x <= lastCenterX + RENDER_RADIUS &&
                update.y >= lastCenterY - RENDER_RADIUS && update.y <= lastCenterY + RENDER_RADIUS)
            {
                Block block = Blocks.AIR;
                if (update.mcName != null && !update.mcName.isEmpty())
                {
                    ResourceLocation loc = new ResourceLocation(update.mcName);
                    Block found = ForgeRegistries.BLOCKS.getValue(loc);
                    if (found != null) block = found;
                }

                level.setBlock(new BlockPos(update.x, update.y, update.z),
                    block == Blocks.AIR ? Blocks.AIR.defaultBlockState() : block.defaultBlockState(), 3);
            }
        }
    }
}
