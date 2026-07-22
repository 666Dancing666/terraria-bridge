using System;
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
        private BridgeEventHandler _eventHandler;

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
            _eventHandler = new BridgeEventHandler(_client);

            _client.OnConnected += OnBridgeConnected;
            _client.OnDisconnected += OnBridgeDisconnected;
            _client.OnMessageReceived += OnMessageReceived;

            ServerApi.Hooks.GameUpdate.Register(this, OnGameUpdate);

            _eventHandler.Register();

            Task.Run(async () => await _client.ConnectAsync());
        }

        protected override void Dispose(bool disposing)
        {
            if (disposing)
            {
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

        private void OnGameUpdate(EventArgs args)
        {
            if (!_client.IsConnected) return;
            _eventHandler.OnUpdate();
        }
    }
}
