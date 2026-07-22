using System;
using System.Collections.Generic;
using Terraria;
using TerrariaBridge.Network;

namespace TerrariaBridge.Handlers
{
    public class BridgeEventHandler
    {
        private BridgeClient _client;
        private int lastDayTime = -1;
        private bool lastBloodMoon = false;
        private bool lastEclipse = false;
        private string lastWeather = "";
        private int weatherTimer = 0;

        public BridgeEventHandler(BridgeClient client)
        {
            _client = client;
        }

        public void Register()
        {
        }

        public void OnUpdate()
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
