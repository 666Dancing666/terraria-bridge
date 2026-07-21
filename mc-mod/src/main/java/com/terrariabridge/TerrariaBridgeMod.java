package com.terrariabridge;

import com.terrariabridge.config.ModConfig;
import com.terrariabridge.network.BridgeClient;
import com.terrariabridge.network.PacketHandler;
import com.terrariabridge.render.WorldRenderer;
import com.terrariabridge.render.LayerManager;
import com.terrariabridge.input.InputInterceptor;
import com.terrariabridge.input.RaycastHandler;
import com.terrariabridge.collision.CollisionDisabler;
import com.terrariabridge.registry.ModBlocks;
import net.minecraftforge.common.MinecraftForge;
import net.minecraftforge.fml.common.Mod;
import net.minecraftforge.fml.event.lifecycle.FMLClientSetupEvent;
import net.minecraftforge.fml.event.lifecycle.FMLCommonSetupEvent;
import net.minecraftforge.fml.javafmlmod.FMLJavaModLoadingContext;

@Mod(TerrariaBridgeMod.MOD_ID)
public class TerrariaBridgeMod
{
    public static final String MOD_ID = "terrariabridge";

    private BridgeClient bridgeClient;
    private PacketHandler packetHandler;
    private WorldRenderer worldRenderer;
    private LayerManager layerManager;
    private InputInterceptor inputInterceptor;
    private RaycastHandler raycastHandler;

    public TerrariaBridgeMod()
    {
        var bus = FMLJavaModLoadingContext.get().getModEventBus();
        bus.addListener(this::onCommonSetup);
        bus.addListener(this::onClientSetup);

        ModBlocks.BLOCKS.register(bus);

        MinecraftForge.EVENT_BUS.register(this);
    }

    public void onCommonSetup(FMLCommonSetupEvent event)
    {
        ModConfig.load();
        packetHandler = new PacketHandler();
        bridgeClient = new BridgeClient(ModConfig.bridgeHost, ModConfig.bridgePort);
        bridgeClient.setOnMessageReceived(packetHandler::handleMessage);
        bridgeClient.connect();
    }

    public void onClientSetup(FMLClientSetupEvent event)
    {
        layerManager = new LayerManager();
        packetHandler.setLayerManager(layerManager);
        worldRenderer = new WorldRenderer(layerManager);
        raycastHandler = new RaycastHandler(layerManager);
        inputInterceptor = new InputInterceptor(raycastHandler, bridgeClient);
        CollisionDisabler.register();

        MinecraftForge.EVENT_BUS.register(inputInterceptor);
        MinecraftForge.EVENT_BUS.register(worldRenderer);
    }
}
