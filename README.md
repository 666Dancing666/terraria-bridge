# Terraria Bridge

让泰拉瑞亚和我的世界玩家实时联机的桥接项目。

注:使用了ai

## 架构

泰拉瑞亚服务端 (TShock) <-> 中间层桥接服务 <-> 我的世界客户端 (Forge Mod)

游戏逻辑以泰拉瑞亚为准，MC 客户端渲染为三层薄片。

## 组件

| 组件 | 目录 | 技术栈 | 状态 |
|------|------|--------|------|
| 中间层服务 | bridge-service/ | Go | 完成 |
| TShock 插件 | tshock-plugin/ | C# (.NET 6) | 代码完成 |
| MC Mod | mc-mod/ | Java (Forge 1.20.1) | 编译通过 |

## 中间层功能

- WebSocket 双向消息转发
- 多 MC 玩家独立会话管理
- 断线重连 + 消息缓冲
- 消息类型过滤
- 消息日志记录
- HTTP 健康检查接口
- 方块/物品/墙壁/实体 ID 映射
- 协议转换 (泰拉瑞亚 <-> Minecraft)

## 版本要求

- 泰拉瑞亚: 1.4.5.6
- TShock: 5.x
- Minecraft: 1.20.1
- Forge: 47.x

## 快速开始

### 中间层服务

cd bridge-service
go run .

TShock 连接: ws://localhost:8080/tshock
MC 连接:     ws://localhost:8080/mc
健康检查:    http://localhost:8080/health

### TShock 插件

cd tshock-plugin
dotnet build

将 DLL 放入 TShock 的 ServerPlugins 目录。

### MC Mod

cd mc-mod
gradle build

将 build/libs 中的 jar 放入 Minecraft 的 mods 目录。

## 坐标映射

泰拉瑞亚 X = MC X
泰拉瑞亚 Y = MC Y
MC Z 分层: 0=前景/实体, 1=家具, 2=背景墙

## 许可证

MIT
