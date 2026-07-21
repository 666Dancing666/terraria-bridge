using System;
using System.IO;
using Newtonsoft.Json;

namespace TerrariaBridge.Config
{
    public class BridgeConfig
    {
        public string BridgeHost { get; set; } = "ws://localhost:8080/tshock";
        public int SyncRadiusX { get; set; } = 100;
        public int SyncRadiusY { get; set; } = 60;
        public int SyncIntervalMs { get; set; } = 50;

        private static string ConfigPath => Path.Combine(
            TShockAPI.TShock.SavePath, "TerrariaBridge.json");

        public static BridgeConfig Load()
        {
            try
            {
                if (File.Exists(ConfigPath))
                {
                    var json = File.ReadAllText(ConfigPath);
                    return JsonConvert.DeserializeObject<BridgeConfig>(json) 
                           ?? new BridgeConfig();
                }
            }
            catch (Exception ex)
            {
                TShockAPI.TShock.Log.ConsoleError(
                    $"TerrariaBridge: 配置加载失败 - {ex.Message}");
            }

            var config = new BridgeConfig();
            config.Save();
            return config;
        }

        public void Save()
        {
            var json = JsonConvert.SerializeObject(this, Formatting.Indented);
            File.WriteAllText(ConfigPath, json);
        }
    }
}
