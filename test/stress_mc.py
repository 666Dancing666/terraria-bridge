import asyncio
import json
import websockets
import time

async def main():
    uri = "ws://localhost:8080/mc"
    async with websockets.connect(uri) as ws:
        print("[MC Stress] Connected")
        count = 0
        start = time.time()
        last = start

        try:
            while True:
                msg = await asyncio.wait_for(ws.recv(), timeout=30)
                count += 1
                now = time.time()
                if now - last >= 1.0:
                    rate = count / (now - start)
                    print(f"[MC Stress] Received {count} msgs, rate: {rate:.1f}/s")
                    last = now
        except asyncio.TimeoutError:
            pass

        total = time.time() - start
        print(f"[MC Stress] Total: {count} msgs in {total:.1f}s, avg: {count/total:.1f}/s")

asyncio.run(main())
