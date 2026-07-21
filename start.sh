#!/bin/bash

echo "========================================"
echo "  Terraria Bridge - 一键启动"
echo "========================================"
echo ""

DIR="$(cd "$(dirname "$0")" && pwd)"

echo "[1/2] 启动中间层服务..."
cd "$DIR/bridge-service"
go build -o bridge-service . 2>/dev/null
./bridge-service &
BRIDGE_PID=$!
echo "  中间层 PID: $BRIDGE_PID"
sleep 1

echo ""
echo "[2/2] 检查服务状态..."
if curl -s http://localhost:8080/health > /dev/null 2>&1; then
    echo "  中间层服务运行正常"
else
    echo "  中间层启动失败，请检查日志"
    exit 1
fi

echo ""
echo "========================================"
echo "  连接信息"
echo "========================================"
echo "  管理面板: http://localhost:8080"
echo "  TShock:   ws://localhost:8080/tshock"
echo "  MC:       ws://localhost:8080/mc"
echo "  健康检查: http://localhost:8080/health"
echo ""
echo "按 Ctrl+C 停止所有服务"
echo "========================================"

trap "echo '正在停止...'; kill $BRIDGE_PID 2>/dev/null; exit 0" INT TERM

wait $BRIDGE_PID
