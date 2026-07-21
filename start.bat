@echo off
echo ========================================
echo   Terraria Bridge - 一键启动 (Windows)
echo ========================================
echo.

cd /d %~dp0

echo [1/2] 编译并启动中间层服务...
cd bridge-service
go build -o bridge-service.exe .
start "TerrariaBridge" bridge-service.exe
cd ..

timeout /t 2 /nobreak > nul

echo [2/2] 检查服务状态...
curl -s http://localhost:8080/health > nul 2>&1
if %errorlevel% equ 0 (
    echo   中间层服务运行正常
) else (
    echo   中间层启动失败
    pause
    exit /b 1
)

echo.
echo ========================================
echo   连接信息
echo ========================================
echo   管理面板: http://localhost:8080
echo   TShock:   ws://localhost:8080/tshock
echo   MC:       ws://localhost:8080/mc
echo ========================================
echo.
pause
