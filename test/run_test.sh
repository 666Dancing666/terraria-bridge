#!/bin/bash
echo "Starting Terraria Bridge integration test..."
echo ""

echo "1. Start the bridge service in another terminal:"
echo "   cd ~/terraria-bridge/bridge-service && go run ."
echo ""
echo "2. Press Enter when bridge is running..."
read

echo ""
echo "3. Starting mock TShock and MC clients..."
echo ""

cd "$(dirname "$0")"

python3 mock_mc.py &
MC_PID=$!
sleep 1

python3 mock_tshock.py &
TS_PID=$!

wait $TS_PID
wait $MC_PID

echo ""
echo "Test complete. Check the bridge service logs."
