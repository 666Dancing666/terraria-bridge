using System;
using System.Net.WebSockets;
using System.Text;
using System.Threading;
using System.Threading.Tasks;
using Newtonsoft.Json;
using TerrariaBridge.Config;

namespace TerrariaBridge.Network
{
    public class BridgeClient
    {
        private ClientWebSocket _ws;
        private CancellationTokenSource _cts;
        private BridgeConfig _config;

        public event Action<BridgeMessage> OnMessageReceived;
        public event Action OnConnected;
        public event Action<string> OnDisconnected;
        public bool IsConnected => _ws?.State == WebSocketState.Open;

        public BridgeClient(BridgeConfig config)
        {
            _config = config;
        }

        public async Task ConnectAsync()
        {
            try
            {
                _ws = new ClientWebSocket();
                _cts = new CancellationTokenSource();

                TShockAPI.TShock.Log.ConsoleInfo(
                    $"TerrariaBridge: 正在连接中间层 {_config.BridgeHost}");

                await _ws.ConnectAsync(
                    new Uri(_config.BridgeHost), _cts.Token);

                TShockAPI.TShock.Log.ConsoleInfo(
                    "TerrariaBridge: 已连接到中间层");

                OnConnected?.Invoke();
                _ = ReceiveLoop();
            }
            catch (Exception ex)
            {
                TShockAPI.TShock.Log.ConsoleError(
                    $"TerrariaBridge: 连接失败 - {ex.Message}");
                OnDisconnected?.Invoke(ex.Message);
            }
        }

        private async Task ReceiveLoop()
        {
            var buffer = new byte[1024 * 64];

            try
            {
                while (_ws.State == WebSocketState.Open)
                {
                    var result = await _ws.ReceiveAsync(
                        new ArraySegment<byte>(buffer), _cts.Token);

                    if (result.MessageType == WebSocketMessageType.Close)
                    {
                        TShockAPI.TShock.Log.ConsoleInfo(
                            "TerrariaBridge: 中间层关闭了连接");
                        break;
                    }

                    var json = Encoding.UTF8.GetString(buffer, 0, result.Count);

                    try
                    {
                        var msg = JsonConvert.DeserializeObject<BridgeMessage>(json);
                        if (msg != null)
                        {
                            OnMessageReceived?.Invoke(msg);
                        }
                    }
                    catch (JsonException ex)
                    {
                        TShockAPI.TShock.Log.ConsoleError(
                            $"TerrariaBridge: JSON解析失败 - {ex.Message}");
                    }
                }
            }
            catch (OperationCanceledException) { }
            catch (Exception ex)
            {
                TShockAPI.TShock.Log.ConsoleError(
                    $"TerrariaBridge: 接收异常 - {ex.Message}");
            }
            finally
            {
                OnDisconnected?.Invoke("连接断开");
            }
        }

        public async Task SendAsync(BridgeMessage message)
        {
            if (!IsConnected) return;

            try
            {
                var json = JsonConvert.SerializeObject(message);
                var bytes = Encoding.UTF8.GetBytes(json);
                await _ws.SendAsync(
                    new ArraySegment<byte>(bytes),
                    WebSocketMessageType.Text,
                    true,
                    _cts.Token);
            }
            catch (Exception ex)
            {
                TShockAPI.TShock.Log.ConsoleError(
                    $"TerrariaBridge: 发送失败 - {ex.Message}");
            }
        }

        public async Task DisconnectAsync()
        {
            _cts?.Cancel();
            if (_ws?.State == WebSocketState.Open)
            {
                await _ws.CloseAsync(
                    WebSocketCloseStatus.NormalClosure, "", CancellationToken.None);
            }
            _ws?.Dispose();
        }
    }
}
