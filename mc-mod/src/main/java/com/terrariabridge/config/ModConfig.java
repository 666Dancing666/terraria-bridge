package com.terrariabridge.config;

import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import java.io.*;
import java.nio.file.*;

public class ModConfig
{
    private static final Gson GSON = new GsonBuilder().setPrettyPrinting().create();
    private static final Path CONFIG_DIR = Path.of("config");
    private static final Path CONFIG_FILE = CONFIG_DIR.resolve("terrariabridge.json");

    public static String bridgeHost = "localhost";
    public static int bridgePort = 8080;

    public static void load()
    {
        try
        {
            Files.createDirectories(CONFIG_DIR);
            if (Files.exists(CONFIG_FILE))
            {
                String json = Files.readString(CONFIG_FILE);
                ConfigData data = GSON.fromJson(json, ConfigData.class);
                if (data != null)
                {
                    bridgeHost = data.bridgeHost;
                    bridgePort = data.bridgePort;
                }
            }
            else
            {
                save();
            }
        }
        catch (IOException e)
        {
        }
    }

    public static void save()
    {
        try
        {
            Files.createDirectories(CONFIG_DIR);
            ConfigData data = new ConfigData();
            data.bridgeHost = bridgeHost;
            data.bridgePort = bridgePort;
            String json = GSON.toJson(data);
            Files.writeString(CONFIG_FILE, json);
        }
        catch (IOException e)
        {
        }
    }

    private static class ConfigData
    {
        String bridgeHost = "localhost";
        int bridgePort = 8080;
    }
}
