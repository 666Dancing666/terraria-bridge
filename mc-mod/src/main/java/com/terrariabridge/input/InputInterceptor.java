package com.terrariabridge.input;

import com.terrariabridge.input.RaycastHandler.TerrariaHitResult;
import com.terrariabridge.network.BridgeClient;
import com.terrariabridge.network.MessageTypes;

import net.minecraft.client.Minecraft;
import net.minecraftforge.client.event.InputEvent;
import net.minecraftforge.event.TickEvent;
import net.minecraftforge.eventbus.api.SubscribeEvent;
import org.lwjgl.glfw.GLFW;

import java.util.HashMap;
import java.util.Map;

public class InputInterceptor
{
    private final RaycastHandler raycastHandler;
    private final BridgeClient bridgeClient;

    private double lastX, lastY;
    private boolean leftClickDown = false;
    private boolean rightClickDown = false;

    public InputInterceptor(RaycastHandler raycastHandler, BridgeClient bridgeClient)
    {
        this.raycastHandler = raycastHandler;
        this.bridgeClient = bridgeClient;
    }

    @SubscribeEvent
    public void onClientTick(TickEvent.ClientTickEvent event)
    {
        if (event.phase != TickEvent.Phase.END) return;

        Minecraft mc = Minecraft.getInstance();
        if (mc.player == null) return;

        mc.player.setPos(mc.player.getX(), mc.player.getY(), 0);

        double currentX = mc.player.getX();
        double currentY = mc.player.getY();

        if (Math.abs(currentX - lastX) > 0.01 || Math.abs(currentY - lastY) > 0.01)
        {
            Map<String, Object> payload = new HashMap<>();
            payload.put("id", mc.player.getId());
            payload.put("x", currentX);
            payload.put("y", currentY);

            String json = MessageTypes.toJson(MessageTypes.PLAYER_MOVE, payload);
            bridgeClient.send(json);

            lastX = currentX;
            lastY = currentY;
        }
    }

    @SubscribeEvent
    public void onMouseClick(InputEvent.MouseButton event)
    {
        Minecraft mc = Minecraft.getInstance();
        if (mc.player == null) return;

        TerrariaHitResult hit = raycastHandler.raycast();
        if (hit == null) return;

        if (event.getButton() == GLFW.GLFW_MOUSE_BUTTON_LEFT &&
            event.getAction() == GLFW.GLFW_PRESS)
        {
            Map<String, Object> payload = new HashMap<>();
            payload.put("x", hit.x);
            payload.put("y", hit.y);
            payload.put("player_id", mc.player.getId());

            String json = MessageTypes.toJson(MessageTypes.TILE_BREAK, payload);
            bridgeClient.send(json);
            event.setCanceled(true);
        }

        if (event.getButton() == GLFW.GLFW_MOUSE_BUTTON_RIGHT &&
            event.getAction() == GLFW.GLFW_PRESS)
        {
            Map<String, Object> payload = new HashMap<>();
            payload.put("x", hit.x);
            payload.put("y", hit.y);
            payload.put("player_id", mc.player.getId());

            String json = MessageTypes.toJson(MessageTypes.INTERACT, payload);
            bridgeClient.send(json);
            event.setCanceled(true);
        }
    }

    @SubscribeEvent
    public void onKeyPress(InputEvent.Key event)
    {
        Minecraft mc = Minecraft.getInstance();
        if (mc.player == null) return;

        if (event.getAction() != GLFW.GLFW_PRESS) return;

        Map<String, Object> payload = new HashMap<>();
        payload.put("player_id", mc.player.getId());

        switch (event.getKey())
        {
            case GLFW.GLFW_KEY_SPACE:
                payload.put("action", "jump");
                String jumpJson = MessageTypes.toJson(MessageTypes.PLAYER_ACTION, payload);
                bridgeClient.send(jumpJson);
                break;
            case GLFW.GLFW_KEY_E:
                break;
        }
    }
}
