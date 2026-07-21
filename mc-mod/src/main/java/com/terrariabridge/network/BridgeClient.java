package com.terrariabridge.network;

import org.java_websocket.client.WebSocketClient;
import org.java_websocket.handshake.ServerHandshake;
import java.net.URI;
import java.util.function.Consumer;

public class BridgeClient
{
    private WebSocketClient client;
    private String host;
    private int port;
    private Consumer<String> onMessageReceived;
    private boolean connected = false;

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
