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
        public override string Description => "连接泰拉瑞亚和我的世界的桥接插件";
        public override Version Version => new Version(0, 1, 0);

        private BridgeClient _client;
        private BridgeConfig _config;
        private WorldSyncHandler _worldSync;
        private EntityHandler _entityHandler;
        private CommandHandler _commandHandler;

        private Dictionary<int, bool> _syncedPlayers = new Dictionary<int, bool>();
        private DateTime _lastEntitySync = DateTime.MinValue;

        public TerrariaBridgePlugin(Main game)
            : base(game)
        {
        }

        public override void Initialize()
        {
            _config = BridgeConfig.Load();
            _client = new BridgeClient(_config);
            _worldSync = new WorldSyncHandler(_client);
            _entityHandler = new EntityHandler(_client);
            _commandHandler = new CommandHandler(_client);

            // 注册事件
            _client.OnConnected += OnBridgeConnected;
            _client.OnDisconnected += OnBridgeDisconnected;
            _client.OnMessageReceived += OnMessageReceived;

            ServerApi.Hooks.NetGreetPlayer.Register(this, OnPlayerJoin);
            ServerApi.Hooks.ServerLeave.Register(this, OnPlayerLeave);
            ServerApi.Hooks.GameUpdate.Register(this, OnGameUpdate);

            // 连接到中间层
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
            TShock.Log.ConsoleInfo("TerrariaBridge: 中间层连接成功，开始同步");
        }

        private void OnBridgeDisconnected(string reason)
        {
            TShock.Log.ConsoleWarn(
                $"TerrariaBridge: 中间层断开 ({reason})，5秒后重连");
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

                // 发送该玩家周围的世界快照
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

            // 限制实体同步频率
            var now = DateTime.UtcNow;
            if ((now - _lastEntitySync).TotalMilliseconds < _config.SyncIntervalMs)
                return;
            _lastEntitySync = now;

            Task.Run(async () =>
            {
                await SyncEntities();
            });
        }

        private async Task SyncEntities()
        {
            try
            {
                // 同步所有活跃玩家
                for (int i = 0; i < Main.player.Length; i++)
                {
                    var player = Main.player[i];
                    if (player != null && player.active)
                    {
                        var msg = _entityHandler.CreatePlayerUpdate(player);
                        await _client.SendAsync(msg);
                    }
                }

                // 同步所有活跃NPC
                for (int i = 0; i < Main.npc.Length; i++)
                {
                    var npc = Main.npc[i];
                    if (npc != null && npc.active)
                    {
                        var msg = _entityHandler.CreateNPCUpdate(npc);
                        await _client.SendAsync(msg);
                    }
                }

                // 同步掉落物品
                for (int i = 0; i < Main.item.Length; i++)
                {
                    var item = Main.item[i];
                    if (item != null && item.active)
                    {
                        var msg = _entityHandler.CreateItemUpdate(item, i);
                        await _client.SendAsync(msg);
                    }
                }

                // 同步投射物
                for (int i = 0; i < Main.projectile.Length; i++)
                {
                    var proj = Main.projectile[i];
                    if (proj != null && proj.active)
                    {
                        var msg = _entityHandler.CreateProjectileUpdate(proj);
                        await _client.SendAsync(msg);
                    }
                }
            }
            catch (Exception ex)
            {
                TShock.Log.ConsoleError($"TerrariaBridge: 实体同步失败 - {ex.Message}");
            }
        }
    }
}
