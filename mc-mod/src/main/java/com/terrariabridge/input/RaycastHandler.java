package com.terrariabridge.input;

import com.terrariabridge.render.LayerManager;
import net.minecraft.client.Minecraft;
import net.minecraft.world.entity.Entity;
import net.minecraft.world.phys.AABB;
import net.minecraft.world.phys.Vec3;
import net.minecraft.core.BlockPos;
import net.minecraft.world.level.Level;
import net.minecraft.world.level.block.Blocks;
import net.minecraft.world.level.block.state.BlockState;
import java.util.List;

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
        public int entityId;

        public TerrariaHitResult(int x, int y, int z, String hitType)
        {
            this.x = x;
            this.y = y;
            this.z = z;
            this.hitType = hitType;
            this.entityId = -1;
        }

        public TerrariaHitResult(int entityId)
        {
            this.entityId = entityId;
            this.hitType = "entity";
        }
    }

    public TerrariaHitResult raycast()
    {
        Minecraft mc = Minecraft.getInstance();
        if (mc.player == null || mc.level == null) return null;

        Level level = mc.level;
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

            BlockState state = level.getBlockState(new BlockPos(bx, by, z));
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

        for (double d = 0.5; d <= range; d += 0.5)
        {
            Vec3 point = eyePos.add(lookVec.scale(d));
            AABB box = new AABB(point.x - 0.3, point.y - 0.3, -1, point.x + 0.3, point.y + 0.3, 2);
            List<Entity> entities = level.getEntitiesOfClass(Entity.class, box, e -> e != mc.player);

            for (Entity entity : entities)
            {
                for (LayerManager.EntityData data : layerManager.getAllEntities())
                {
                    double ex = data.x / 16.0;
                    double ey = data.y / 16.0;

                    if (Math.abs(ex - entity.getX()) < 0.5 && Math.abs(ey - entity.getY()) < 0.5)
                    {
                        return new TerrariaHitResult(data.id);
                    }
                }
            }
        }

        return null;
    }
}
