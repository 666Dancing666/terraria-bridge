package com.terrariabridge.collision;

import net.minecraft.world.entity.Entity;
import net.minecraft.world.entity.player.Player;
import net.minecraft.world.entity.Mob;
import net.minecraftforge.event.entity.living.LivingEvent;
import net.minecraftforge.eventbus.api.SubscribeEvent;
import net.minecraftforge.common.MinecraftForge;

public class CollisionDisabler
{
    public static void register()
    {
        MinecraftForge.EVENT_BUS.register(new CollisionDisabler());
    }

    @SubscribeEvent
    public void onLivingTick(LivingEvent.LivingTickEvent event)
    {
        Entity entity = event.getEntity();

        if (entity instanceof Player)
        {
            entity.noPhysics = false;
            entity.setOnGround(true);
        }

        if (entity instanceof Mob)
        {
            entity.noPhysics = false;
        }
    }
}
