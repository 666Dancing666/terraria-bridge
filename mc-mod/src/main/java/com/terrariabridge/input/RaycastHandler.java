package com.terrariabridge.input;

import com.terrariabridge.render.LayerManager;
import net.minecraft.client.Minecraft;
import net.minecraft.world.phys.HitResult;
import net.minecraft.world.phys.BlockHitResult;
import net.minecraft.world.phys.Vec3;
import net.minecraft.core.BlockPos;
import net.minecraft.world.level.ClipContext;
import net.minecraft.world.level.Level;
import net.minecraft.world.level.block.Blocks;
import net.minecraft.world.level.block.state.BlockState;

public class RaycastHandler
{
    private final LayerManager layerManager;

    public RaycastHandler(LayerManager layerManager)
    {
        this.layerManager = layerManager;
    }

    public static class TerrariaHitResult
    {
        public int x, y, z;
        public String hitType;

        public TerrariaHitResult(int x, int y, int z, String hitType)
        {
            this.x = x;
            this.y = y;
            this.z = z;
            this.hitType = hitType;
        }
    }

    public TerrariaHitResult raycast()
    {
        Minecraft mc = Minecraft.getInstance();
        if (mc.player == null || mc.level == null) return null;

        Vec3 eyePos = mc.player.getEyePosition();
        Vec3 lookVec = mc.player.getLookAngle();
        double range = 20.0;

        int[] layers = {
            LayerManager.LAYER_FOREGROUND,
            LayerManager.LAYER_FURNITURE,
            LayerManager.LAYER_WALL
        };

        for (int z : layers)
        {
            double dz = z + 0.5 - eyePos.z;
            if (Math.abs(lookVec.z) < 0.0001) continue;

            double t = dz / lookVec.z;
            if (t < 0 || t > range) continue;

            double hitX = eyePos.x + lookVec.x * t;
            double hitY = eyePos.y + lookVec.y * t;

            int bx = (int) Math.floor(hitX);
            int by = (int) Math.floor(hitY);

            BlockState state = mc.level.getBlockState(new BlockPos(bx, by, z));
            if (state.getBlock() != Blocks.AIR)
            {
                String hitType;
                switch (z)
                {
                    case LayerManager.LAYER_FOREGROUND:
                        hitType = "tile";
                        break;
                    case LayerManager.LAYER_FURNITURE:
                        hitType = "furniture";
                        break;
                    case LayerManager.LAYER_WALL:
                        hitType = "wall";
                        break;
                    default:
                        hitType = "unknown";
                }

                return new TerrariaHitResult(bx, by, z, hitType);
            }
        }

        return null;
    }
}
