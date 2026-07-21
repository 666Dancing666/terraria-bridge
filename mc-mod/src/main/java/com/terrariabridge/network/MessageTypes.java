package com.terrariabridge.network;

import com.google.gson.*;
import java.util.Map;

public class MessageTypes
{
    private static final Gson GSON = new Gson();

    public static class BridgeMessage
    {
        public String type;
        public Map<String, Object> payload;
    }

    public static final String WORLD_SNAPSHOT = "world_snapshot";
    public static final String TILE_UPDATE = "tile_update";
    public static final String WALL_UPDATE = "wall_update";
    public static final String LIQUID_UPDATE = "liquid_update";
    public static final String ENTITY_UPDATE = "entity_update";
    public static final String ENTITY_REMOVE = "entity_remove";
    public static final String PLAYER_JOIN = "player_join";
    public static final String PLAYER_LEAVE = "player_leave";
    public static final String CHAT_MESSAGE = "chat_message";

    public static final String PLAYER_MOVE = "player_move";
    public static final String PLAYER_ACTION = "player_action";
    public static final String TILE_BREAK = "tile_break";
    public static final String TILE_PLACE = "tile_place";
    public static final String INTERACT = "interact";

    public static BridgeMessage parse(String json)
    {
        try
        {
            return GSON.fromJson(json, BridgeMessage.class);
        }
        catch (JsonSyntaxException e)
        {
            return null;
        }
    }

    public static String toJson(String type, Map<String, Object> payload)
    {
        BridgeMessage msg = new BridgeMessage();
        msg.type = type;
        msg.payload = payload;
        return GSON.toJson(msg);
    }
}
