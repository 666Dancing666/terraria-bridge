package com.terrariabridge.registry;

import net.minecraft.world.level.block.Block;
import net.minecraft.world.level.block.SoundType;
import net.minecraft.world.level.block.state.BlockBehaviour;
import net.minecraft.world.level.material.MapColor;
import net.minecraftforge.registries.DeferredRegister;
import net.minecraftforge.registries.ForgeRegistries;
import net.minecraftforge.registries.RegistryObject;

public class ModBlocks {
    public static final DeferredRegister<Block> BLOCKS =
        DeferredRegister.create(ForgeRegistries.BLOCKS, "terrariabridge");

    public static final RegistryObject<Block> UNKNOWN_BLOCK = BLOCKS.register("unknown_block",
        () -> new Block(BlockBehaviour.Properties.of()
            .mapColor(MapColor.COLOR_PURPLE)
            .strength(1.0f)
            .sound(SoundType.STONE)));

    public static final RegistryObject<Block> UNKNOWN_WALL = BLOCKS.register("unknown_wall",
        () -> new Block(BlockBehaviour.Properties.of()
            .mapColor(MapColor.COLOR_MAGENTA)
            .strength(1.0f)
            .sound(SoundType.WOOL)));
}
