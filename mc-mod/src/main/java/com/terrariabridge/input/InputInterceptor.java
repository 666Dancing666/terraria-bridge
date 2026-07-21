package com.terrariabridge.input;

import com.terrariabridge.input.RaycastHandler.TerrariaHitResult;
import com.terrariabridge.network.BridgeClient;
import com.terrariabridge.network.MessageTypes;
import com.terrariabridge.gui.RecipeScreen;
import com.google.gson.Gson;
import com.google.gson.reflect.TypeToken;

import net.minecraft.client.Minecraft;
import net.minecraft.world.entity.player.Player;
import net.minecraftforge.client.event.InputEvent;
import net.minecraftforge.event.TickEvent;
import net.minecraftforge.eventbus.api.SubscribeEvent;
import org.lwjgl.glfw.GLFW;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class InputInterceptor
{
    private final RaycastHandler raycastHandler;
    private final BridgeClient bridgeClient;
    private double lastX, lastY;
    private boolean leftWasDown = false;
    private boolean rightWasDown = false;
    private boolean jumpWasDown = false;
    private boolean rWasDown = false;
    private Gson gson = new Gson();

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
        if (mc.player == null || mc.level == null) return;

        Player player = mc.player;
        long window = mc.getWindow().getWindow();

        double currentX = player.getX();
        double currentY = player.getY();
        player.setPos(currentX, currentY, 0.0);

        if (Math.abs(currentX - lastX) > 0.05 || Math.abs(currentY - lastY) > 0.05)
        {
            Map<String, Object> payload = new HashMap<>();
            payload.put("id", player.getId());
            payload.put("x", currentX);
            payload.put("y", currentY);
            payload.put("z", 0);
            bridgeClient.send(MessageTypes.toJson(MessageTypes.PLAYER_MOVE, payload));
            lastX = currentX;
            lastY = currentY;
        }

        boolean spaceDown = GLFW.glfwGetKey(window, GLFW.GLFW_KEY_SPACE) == GLFW.GLFW_PRESS;
        if (spaceDown && !jumpWasDown)
        {
            Map<String, Object> payload = new HashMap<>();
            payload.put("player_id", player.getId());
            payload.put("action", "jump");
            bridgeClient.send(MessageTypes.toJson(MessageTypes.PLAYER_ACTION, payload));
        }
        jumpWasDown = spaceDown;

        boolean rDown = GLFW.glfwGetKey(window, GLFW.GLFW_KEY_R) == GLFW.GLFW_PRESS;
        if (rDown && !rWasDown)
        {
            if (mc.screen == null)
            {
                bridgeClient.send("{\"type\":\"request_recipes\"}");
            }
            else if (mc.screen instanceof RecipeScreen)
            {
                mc.setScreen(null);
            }
        }
        rWasDown = rDown;

        boolean leftDown = GLFW.glfwGetMouseButton(window, GLFW.GLFW_MOUSE_BUTTON_LEFT) == GLFW.GLFW_PRESS;
        boolean rightDown = GLFW.glfwGetMouseButton(window, GLFW.GLFW_MOUSE_BUTTON_RIGHT) == GLFW.GLFW_PRESS;

        TerrariaHitResult hit = raycastHandler.raycast();

        if (leftDown && !leftWasDown && hit != null)
        {
            Map<String, Object> payload = new HashMap<>();
            payload.put("x", hit.x);
            payload.put("y", hit.y);
            payload.put("player_id", player.getId());
            bridgeClient.send(MessageTypes.toJson(MessageTypes.TILE_BREAK, payload));
        }

        if (rightDown && !rightWasDown && hit != null)
        {
            Map<String, Object> payload = new HashMap<>();
            payload.put("x", hit.x);
            payload.put("y", hit.y);
            payload.put("player_id", player.getId());
            bridgeClient.send(MessageTypes.toJson(MessageTypes.INTERACT, payload));
        }

        leftWasDown = leftDown;
        rightWasDown = rightDown;
    }

    public static void openRecipeScreen(List<Map<String, Object>> recipes, BridgeClient client)
    {
        Minecraft mc = Minecraft.getInstance();
        RecipeScreen screen = new RecipeScreen(client);
        screen.setRecipes(recipes);
        mc.setScreen(screen);
    }
}
