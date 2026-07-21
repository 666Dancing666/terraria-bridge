package com.terrariabridge.input;

import com.terrariabridge.render.LayerManager;
import net.minecraft.client.Minecraft;
import net.minecraft.world.phys.HitResult;
import net.minecraft.world.phys.BlockHitResult;
import net.minecraft.world.phys.Vec3;
import net.minecraft.core.BlockPos;
import net.minecraft.world.level.ClipContext;
import net.minecraft.world.level.Level;

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
        public String hitType; // "tile", "furniture", "wall", "entity"
        public int entityId;

        public TerrariaHitResult(int x, int y, int z)
        {
            this.x = x;
            this.y = y;
            this.z = z;
        }
    }

    public TerrariaHitResult raycast()
    {
        Minecraft mc = Minecraft.getInstance();
        if (mc.player == null || mc.level == null) return null;

        Vec3 eyePos = mc.player.getEyePosition();
        Vec3 lookVec = mc.player.getLookAngle();
        Vec3 endPos = eyePos.add(lookVec.scale(20.0));

        for (int z = 0; z <= 2; z++)
        {
            ClipContext context = new ClipContext(
                eyePos, endPos,
                ClipContext.Block.OUTLINE,
                ClipContext.Fluid.NONE,
                mc.player
            );

            BlockHitResult hit = mc.level.clip(context);

            if (hit.getType() == HitResult.Type.BLOCK)
            {
                BlockPos pos = hit.getBlockPos();
                if (pos.getZ() == z)
                {
                    TerrariaHitResult result = new TerrariaHitResult(
                        pos.getX(), pos.getY(), pos.getZ());

                    switch (z)
                    {
                        case LayerManager.LAYER_FOREGROUND:
                            result.hitType = "tile";
                            break;
                        case LayerManager.LAYER_FURNITURE:
                            result.hitType = "furniture";
                            break;
                        case LayerManager.LAYER_WALL:
                            result.hitType = "wall";
                            break;
                    }

                    return result;
                }
            }
        }

        return null;
    }
}
