# Terraria Bridge

让泰拉瑞亚和我的世界玩家实时联机的桥接项目。

## 架构

泰拉瑞亚 (TShock) <-> 中间层桥接服务 <-> 我的世界 (Forge Mod)

游戏逻辑以泰拉瑞亚为准，MC 渲染为三层薄片 (Z=0 前景/实体, Z=1 家具, Z=2 背景墙)。

## 功能

- 世界同步：MC 中实时看到泰拉瑞亚世界
- 玩家操作：移动、跳跃、挖掘、放置、交互
- 实体同步：NPC、怪物、掉落物、投射物
- 聊天互通：TShock <-> MC 双向聊天
- 时间同步：泰拉瑞亚时间 <-> MC 时间
- 天气同步：泰拉瑞亚天气 <-> MC 天气
- 事件通知：Boss、血月、日食
- 配方合成：按 R 打开合成界面
- 钩子：按 F 使用钩子
- Web 管理面板：浏览器查看在线状态和延迟

## 版本要求

- 泰拉瑞亚: 1.4.5.6
- TShock: 6.1.0
- Minecraft: 1.20.1
- Forge: 47.x
- Go: 1.21+ (仅编译中间层时需要)
- .NET: 9.0 (仅编译 TShock 插件时需要)
- JDK: 17 (仅编译 MC Mod 时需要)

## 快速开始

### 1. 下载

从 Release 页面下载最新版本，解压后得到：

- bridge-service (中间层)
- TerrariaBridge.dll (TShock 插件)
- terraria-bridge-mc-mod-*.jar (MC Mod)

### 2. 启动中间层

Linux/Mac:
./bridge-service

Windows:
bridge-service.exe

启动后：
- 管理面板: http://localhost:8080
- TShock 连接: ws://localhost:8080/tshock
- MC 连接: ws://localhost:8080/mc
- 健康检查: http://localhost:8080/health

### 3. 安装 TShock 插件

将 TerrariaBridge.dll 复制到 TShock 的 ServerPlugins 目录，启动 TShock 即可自动连接中间层。

### 4. 安装 MC Mod

将 jar 文件复制到 Minecraft 的 mods 目录，启动 Minecraft (Forge 1.20.1)，进入世界即可自动连接中间层。

### 5. 操作说明

- 移动: WASD
- 跳跃: 空格
- 挖掘: 鼠标左键
- 交互/放置: 鼠标右键
- 合成: R 键
- 钩子: F 键
- 聊天: T 键

## 目录结构

bridge-service/   中间层源码 (Go)
tshock-plugin/    TShock 插件源码 (C#)
mc-mod/           MC Mod 源码 (Java)
test/             测试脚本 (Python)
config/           映射表

## 许可证

MIT
