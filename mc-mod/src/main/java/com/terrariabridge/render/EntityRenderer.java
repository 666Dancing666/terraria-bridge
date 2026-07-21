package com.terrariabridge.render;

import net.minecraft.client.Minecraft;
import net.minecraft.world.entity.Entity;
import net.minecraft.world.entity.EntityType;
import net.minecraft.world.entity.decoration.ArmorStand;
import net.minecraft.world.entity.monster.Zombie;
import net.minecraft.world.entity.monster.Skeleton;
import net.minecraft.world.entity.monster.Slime;
import net.minecraft.world.entity.monster.Blaze;
import net.minecraft.world.entity.projectile.Arrow;
import net.minecraft.world.entity.projectile.Fireball;
import net.minecraft.world.entity.item.ItemEntity;
import net.minecraft.world.item.ItemStack;
import net.minecraft.world.item.Items;
import net.minecraft.resources.ResourceLocation;
import net.minecraftforge.registries.ForgeRegistries;
import net.minecraft.world.level.Level;
import net.minecraftforge.event.TickEvent;
import net.minecraftforge.eventbus.api.SubscribeEvent;
import java.util.*;

public class EntityRenderer
{
    private final LayerManager layerManager;
    private final Map<Integer, Entity> spawnedEntities = new HashMap<>();
    private final Set<Integer> knownEntityIds = new HashSet<>();

    public EntityRenderer(LayerManager layerManager)
    {
        this.layerManager = layerManager;
    }

    @SubscribeEvent
    public void onClientTick(TickEvent.ClientTickEvent event)
    {
        if (event.phase != TickEvent.Phase.END) return;

        Minecraft mc = Minecraft.getInstance();
        if (mc.player == null || mc.level == null) return;

        Level level = mc.level;
        Collection<LayerManager.EntityData> entities = layerManager.getAllEntities();

        Set<Integer> currentIds = new HashSet<>();

        for (LayerManager.EntityData data : entities)
        {
            currentIds.add(data.id);

            Entity entity = spawnedEntities.get(data.id);
            if (entity == null || entity.isRemoved())
            {
                entity = createEntity(level, data);
                if (entity != null)
                {
                    level.addFreshEntity(entity);
                    spawnedEntities.put(data.id, entity);
                }
            }

            if (entity != null)
            {
                double mcX = data.x / 16.0;
                double mcY = data.y / 16.0;
                double mcZ = 0.5;

                entity.setPos(mcX, mcY, mcZ);
                entity.setDeltaMovement(0, 0, 0);

                if (data.velocityX != 0 || data.velocityY != 0)
                {
                    entity.setDeltaMovement(
                        data.velocityX / 16.0,
                        data.velocityY / 16.0,
                        0);
                }

                if (entity instanceof ArmorStand stand)
                {
                    stand.setCustomNameVisible(true);
                    if (data.name != null && !data.name.isEmpty())
                    {
                        stand.setCustomName(
                            net.minecraft.network.chat.Component.literal(data.name));
                    }
                }
            }
        }

        for (int id : new HashSet<>(spawnedEntities.keySet()))
        {
            if (!currentIds.contains(id))
            {
                Entity entity = spawnedEntities.get(id);
                if (entity != null)
                {
                    entity.remove(Entity.RemovalReason.DISCARDED);
                }
                spawnedEntities.remove(id);
            }
        }

        knownEntityIds.clear();
        knownEntityIds.addAll(currentIds);
    }

    private Entity createEntity(Level level, LayerManager.EntityData data)
    {
        String entityType = data.entityType;
        String mcEntity = data.mcEntity;

        if (mcEntity == null)
        {
            mcEntity = "minecraft:armor_stand";
        }

        double x = data.x / 16.0;
        double y = data.y / 16.0;
        double z = 0.5;

        if (entityType != null && entityType.equals("projectile"))
        {
            return createProjectile(level, data, x, y, z);
        }

        switch (mcEntity)
        {
            case "minecraft:zombie":
            case "minecraft:husk":
            case "minecraft:drowned":
                Zombie zombie = new Zombie(EntityType.ZOMBIE, level);
                zombie.setPos(x, y, z);
                zombie.setNoAi(true);
                zombie.setSilent(true);
                return zombie;

            case "minecraft:skeleton":
            case "minecraft:wither_skeleton":
                Skeleton skeleton = new Skeleton(EntityType.SKELETON, level);
                skeleton.setPos(x, y, z);
                skeleton.setNoAi(true);
                skeleton.setSilent(true);
                return skeleton;

            case "minecraft:slime":
            case "minecraft:magma_cube":
                Slime slime = new Slime(EntityType.SLIME, level);
                slime.setPos(x, y, z);
                slime.setNoAi(true);
                slime.setSilent(true);
                return slime;

            case "minecraft:blaze":
                Blaze blaze = new Blaze(EntityType.BLAZE, level);
                blaze.setPos(x, y, z);
                blaze.setNoAi(true);
                blaze.setSilent(true);
                return blaze;

            case "minecraft:player":
                ArmorStand playerStand = new ArmorStand(EntityType.ARMOR_STAND, level);
                playerStand.setPos(x, y, z);
                playerStand.setInvisible(false);
                playerStand.setNoGravity(true);
                playerStand.setNoBasePlate(true);
                if (data.name != null)
                {
                    playerStand.setCustomName(
                        net.minecraft.network.chat.Component.literal(data.name));
                    playerStand.setCustomNameVisible(true);
                }
                return playerStand;

            default:
                if (entityType != null && entityType.equals("item"))
                {
                    ItemEntity item = new ItemEntity(EntityType.ITEM, level);
                    item.setPos(x, y, z);
                    if (data.mcItem != null) {
                    net.minecraft.world.item.Item mcItem = ForgeRegistries.ITEMS.getValue(new ResourceLocation(data.mcItem));
                    if (mcItem != null) item.setItem(new ItemStack(mcItem));
                    else item.setItem(new ItemStack(Items.DIAMOND));
                } else {
                    item.setItem(new ItemStack(Items.DIAMOND));
                }
                    item.setNoPickUpDelay();
                    item.setNeverPickUp();
                    return item;
                }

                ArmorStand stand = new ArmorStand(EntityType.ARMOR_STAND, level);
                stand.setPos(x, y, z);
                stand.setInvisible(false);
                stand.setNoGravity(true);
                stand.setNoBasePlate(true);
                if (data.name != null)
                {
                    stand.setCustomName(
                        net.minecraft.network.chat.Component.literal(data.name));
                    stand.setCustomNameVisible(true);
                }
                return stand;
        }
    }

    private Entity createProjectile(Level level, LayerManager.EntityData data, double x, double y, double z)
    {
        String mcProj = data.mcEntity;
        if (mcProj == null) mcProj = "minecraft:arrow";

        switch (mcProj)
        {
            case "minecraft:arrow":
                Arrow arrow = new Arrow(EntityType.ARROW, level);
                arrow.setPos(x, y, z);
                arrow.setNoGravity(true);
                arrow.pickup = Arrow.Pickup.DISALLOWED;
                return arrow;

            case "minecraft:fireball":
                Fireball fireball = new net.minecraft.world.entity.projectile.SmallFireball(EntityType.SMALL_FIREBALL, level);
                fireball.setPos(x, y, z);
                return fireball;

            default:
                Arrow defaultArrow = new Arrow(EntityType.ARROW, level);
                defaultArrow.setPos(x, y, z);
                defaultArrow.setNoGravity(true);
                defaultArrow.pickup = Arrow.Pickup.DISALLOWED;
                return defaultArrow;
        }
    }
}
