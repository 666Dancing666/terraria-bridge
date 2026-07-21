using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Terraria;
using TerrariaApi.Server;
using TShockAPI;
using TerrariaBridge.Config;
using TerrariaBridge.Network;
using TerrariaBridge.Handlers;

namespace TerrariaBridge
{
    [ApiVersion(2, 1)]
    public class TerrariaBridgePlugin : TerrariaPlugin
    {
        public override string Name => "TerrariaBridge";
        public override string Author => "TerrariaBridge Team";
        public override string Description => "Terraria-Minecraft bridge plugin";
        public override Version Version => new Version(0, 1, 0);

        public static TerrariaBridgePlugin Instance { get; private set; }

        private BridgeClient _client;
        private BridgeConfig _config;
        private WorldSyncHandler _worldSync;
        private EntityHandler _entityHandler;
        private CommandHandler _commandHandler;
        private EventHandler _eventHandler;
        private TileUpdateHandler _tileUpdateHandler;

        public TerrariaBridgePlugin(Main game) : base(game)
        {
            Instance = this;
        }

        public override void Initialize()
        {
            _config = BridgeConfig.Load();
            _client = new BridgeClient(_config);
            _worldSync = new WorldSyncHandler(_client);
            _entityHandler = new EntityHandler(_client);
            _commandHandler = new CommandHandler(_client);
            _eventHandler = new EventHandler(_client);
            _tileUpdateHandler = new TileUpdateHandler(_client, _worldSync);

            _client.OnConnected += OnBridgeConnected;
            _client.OnDisconnected += OnBridgeDisconnected;
            _client.OnMessageReceived += OnMessageReceived;

            ServerApi.Hooks.NetGreetPlayer.Register(this, OnPlayerJoin);
            ServerApi.Hooks.ServerLeave.Register(this, OnPlayerLeave);
            ServerApi.Hooks.GameUpdate.Register(this, OnGameUpdate);

            _eventHandler.Register();
            _tileUpdateHandler.Register();

            Task.Run(async () => await _client.ConnectAsync());
        }

        protected override void Dispose(bool disposing)
        {
            if (disposing)
            {
                ServerApi.Hooks.NetGreetPlayer.Deregister(this, OnPlayerJoin);
                ServerApi.Hooks.ServerLeave.Deregister(this, OnPlayerLeave);
                ServerApi.Hooks.GameUpdate.Deregister(this, OnGameUpdate);

                Task.Run(async () => await _client.DisconnectAsync());
            }
            base.Dispose(disposing);
        }

        private void OnBridgeConnected()
        {
            TShock.Log.ConsoleInfo("TerrariaBridge: Connected to bridge");
        }

        private void OnBridgeDisconnected(string reason)
        {
            TShock.Log.ConsoleWarn($"TerrariaBridge: Disconnected ({reason}), retrying in 5s");
            Task.Run(async () =>
            {
                await Task.Delay(5000);
                await _client.ConnectAsync();
            });
        }

        private void OnMessageReceived(BridgeMessage msg)
        {
            _commandHandler.HandleMessage(msg);
        }

        private void OnPlayerJoin(NetGreetPlayerEventArgs args)
        {
            var player = TShock.Players[args.Who];
            if (player == null) return;

            Task.Run(async () =>
            {
                var joinMsg = _entityHandler.CreatePlayerJoin(player.TPlayer);
                await _client.SendAsync(joinMsg);

                var snapshot = _worldSync.CreateFullSnapshot(
                    (int)(player.TPlayer.position.X / 16),
                    (int)(player.TPlayer.position.Y / 16),
                    _config.SyncRadiusX,
                    _config.SyncRadiusY);
                await _client.SendAsync(snapshot);
            });
        }

        private void OnPlayerLeave(ServerLeaveEventArgs args)
        {
            Task.Run(async () =>
            {
                var leaveMsg = _entityHandler.CreatePlayerLeave(args.Who);
                await _client.SendAsync(leaveMsg);
            });
        }

        private void OnGameUpdate(GameUpdateEventArgs args)
        {
            if (!_client.IsConnected) return;
        }
    }
}
