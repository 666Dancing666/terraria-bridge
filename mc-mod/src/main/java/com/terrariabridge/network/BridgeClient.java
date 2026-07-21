package com.terrariabridge.network;

import net.minecraft.client.Minecraft;
import net.minecraft.network.chat.Component;
import org.java_websocket.client.WebSocketClient;
import org.java_websocket.handshake.ServerHandshake;
import com.google.gson.Gson;
import com.google.gson.JsonObject;
import com.google.gson.JsonArray;
import com.google.gson.reflect.TypeToken;
import com.terrariabridge.input.InputInterceptor;
import java.net.URI;
import java.util.List;
import java.util.Map;
import java.util.function.Consumer;

public class BridgeClient
{
    private WebSocketClient client;
    private String host;
    private int port;
    private Consumer<String> onMessageReceived;
    private boolean connected = false;
    private Gson gson = new Gson();

    public BridgeClient(String host, int port)
    {
        this.host = host;
        this.port = port;
    }

    public void setOnMessageReceived(Consumer<String> callback)
    {
        this.onMessageReceived = callback;
    }

    public boolean isConnected()
    {
        return connected;
    }

    public void connect()
    {
        try
        {
            URI uri = new URI("ws://" + host + ":" + port + "/mc");
            client = new WebSocketClient(uri)
            {
                @Override
                public void onOpen(ServerHandshake handshake)
                {
                    System.out.println("TerrariaBridge: Connected to bridge");
                    connected = true;
                }

                @Override
                public void onMessage(String message)
                {
                    try
                    {
                        JsonObject obj = gson.fromJson(message, JsonObject.class);
                        String type = obj.get("type").getAsString();

                        if (type.equals("chat_message") && obj.has("payload"))
                        {
                            JsonObject payload = obj.getAsJsonObject("payload");
                            String chatMsg = payload.get("msg").getAsString();
                            Minecraft mc = Minecraft.getInstance();
                            if (mc.player != null)
                            {
                                mc.player.displayClientMessage(
                                    Component.literal(chatMsg), false);
                            }
                        }

                        if (type.equals("recipe_list") && obj.has("payload"))
                        {
                            JsonObject payload = obj.getAsJsonObject("payload");
                            JsonArray arr = payload.getAsJsonArray("recipes");
                            List<Map<String, Object>> recipes = gson.fromJson(
                                arr, new TypeToken<List<Map<String, Object>>>(){}.getType());
                            InputInterceptor.openRecipeScreen(recipes, BridgeClient.this);
                        }
                    }
                    catch (Exception e)
                    {
                        System.err.println("Parse error: " + e.getMessage());
                    }

                    if (onMessageReceived != null)
                    {
                        onMessageReceived.accept(message);
                    }
                }

                @Override
                public void onClose(int code, String reason, boolean remote)
                {
                    System.out.println("TerrariaBridge: Disconnected - " + reason);
                    connected = false;
                    new Thread(() ->
                    {
                        try { Thread.sleep(5000); }
                        catch (InterruptedException e) {}
                        connect();
                    }).start();
                }

                @Override
                public void onError(Exception ex)
                {
                    System.err.println("TerrariaBridge: Error - " + ex.getMessage());
                    connected = false;
                }
            };
            client.connect();
        }
        catch (Exception e)
        {
            System.err.println("TerrariaBridge: Connection failed - " + e.getMessage());
        }
    }

    public void send(String message)
    {
        if (client != null && client.isOpen())
        {
            client.send(message);
        }
    }

    public void disconnect()
    {
        if (client != null)
        {
            client.close();
        }
    }
}
