using System;
using System.Collections.Generic;
using Terraria;
using TerrariaApi.Server;
using TShockAPI;
using TerrariaBridge.Network;

namespace TerrariaBridge.Handlers
{
    public class EventHandler
    {
        private BridgeClient _client;

        public EventHandler(BridgeClient client)
        {
            _client = client;
        }

        public void Register()
        {
            ServerApi.Hooks.NpcSpawn.Register(TerrariaBridgePlugin.Instance, OnNpcSpawn);
            ServerApi.Hooks.NpcKilled.Register(TerrariaBridgePlugin.Instance, OnNpcKilled);
            ServerApi.Hooks.PlayerDeath.Register(TerrariaBridgePlugin.Instance, OnPlayerDeath);
            ServerApi.Hooks.GameUpdate.Register(TerrariaBridgePlugin.Instance, OnGameUpdate);
        }

        private int lastDayTime = -1;
        private bool lastBloodMoon = false;
        private bool lastEclipse = false;
        private string lastWeather = "";
        private int weatherTimer = 0;

        private void OnNpcSpawn(NpcSpawnEventArgs args)
        {
            int npcIndex = args.NpcId;
            if (npcIndex < 0 || npcIndex >= Main.npc.Length) return;

            NPC npc = Main.npc[npcIndex];
            if (npc == null || !npc.active) return;

            if (npc.boss)
            {
                SendEvent("boss_spawn", npc.FullName);
            }
        }

        private void OnNpcKilled(NpcKilledEventArgs args)
        {
            int npcIndex = args.npcId;
            if (npcIndex < 0 || npcIndex >= Main.npc.Length) return;

            NPC npc = Main.npc[npcIndex];
            if (npc == null) return;

            if (npc.boss)
            {
                SendEvent("boss_death", npc.FullName);
            }
        }

        private void OnGameUpdate(GameUpdateEventArgs args)
        {
            if (!_client.IsConnected) return;

            int currentTime = (int)Main.time;
            bool isDay = Main.dayTime;

            if (lastDayTime != currentTime / 3600)
            {
                lastDayTime = currentTime / 3600;

                var payload = new Dictionary<string, object>
                {
                    { "time", Main.time },
                    { "is_day", isDay },
                    { "day_time", Main.dayTime }
                };

                _client.SendAsync(new BridgeMessage
                {
                    Type = "time_sync",
                    Payload = payload
                }).Wait();
            }

            weatherTimer++;
            if (weatherTimer >= 120)
            {
                weatherTimer = 0;

                string weather = "clear";
                string wind = "calm";

                if (Main.raining)
                {
                    weather = Main.maxRaining > 0.6f ? "heavy_rain" : "rain";
                }
                if (Main.sandTiles > 100)
                {
                    weather = "sandstorm";
                }
                if (Main.windSpeedCurrent > 0.5f)
                {
                    wind = "strong";
                }

                string currentWeather = weather + "|" + wind;
                if (currentWeather != lastWeather)
                {
                    lastWeather = currentWeather;

                    var payload = new Dictionary<string, object>
                    {
                        { "weather", weather },
                        { "wind", wind }
                    };

                    _client.SendAsync(new BridgeMessage
                    {
                        Type = "weather_sync",
                        Payload = payload
                    }).Wait();
                }
            }

            if (Main.bloodMoon != lastBloodMoon)
            {
                lastBloodMoon = Main.bloodMoon;
                SendEvent(lastBloodMoon ? "blood_moon_start" : "blood_moon_end", "");
            }

            if (Main.eclipse != lastEclipse)
            {
                lastEclipse = Main.eclipse;
                SendEvent(lastEclipse ? "eclipse_start" : "eclipse_end", "");
            }
        }

        private void OnPlayerDeath(PlayerDeathEventArgs args)
        {
            var player = TShock.Players[args.PlayerId];
            if (player != null)
            {
                SendEvent("player_death", player.Name);
            }
        }

        private void SendEvent(string eventType, string eventName)
        {
            var payload = new Dictionary<string, object>
            {
                { "event_type", eventType },
                { "event_name", eventName }
            };

            _client.SendAsync(new BridgeMessage
            {
                Type = "game_event",
                Payload = payload
            }).Wait();
        }
    }
}
