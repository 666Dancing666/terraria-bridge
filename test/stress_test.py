import asyncio
import json
import websockets
import time
import random

async def main():
    uri = "ws://localhost:8080/tshock"
    async with websockets.connect(uri) as ws:
        print("[Stress] Connected")

        start = time.time()

        tiles = []
        for x in range(100):
            for y in range(50):
                tiles.append({
                    "x": x, "y": y,
                    "tile_type": random.randint(1, 100),
                    "wall_type": random.randint(0, 5)
                })

        snapshot = {
            "type": "world_snapshot",
            "payload": {
                "center_x": 50, "center_y": 25,
                "radius_x": 100, "radius_y": 50,
                "world_width": 200, "world_height": 100,
                "tiles": tiles
            }
        }

        t1 = time.time()
        await ws.send(json.dumps(snapshot))
        t2 = time.time()
        print(f"[Stress] Sent 5000 tiles in {(t2-t1)*1000:.1f}ms")

        await asyncio.sleep(0.5)

        for batch in range(10):
            t1 = time.time()
            for i in range(100):
                npc = {
                    "type": "entity_update",
                    "payload": {
                        "entity_type": "npc",
                        "id": batch * 100 + i,
                        "net_id": random.choice([1, 3, 6, 7, 21, 26, 34, 43]),
                        "x": random.randint(0, 200) * 16,
                        "y": random.randint(0, 100) * 16,
                        "health": random.randint(50, 500),
                        "max_health": 500,
                        "direction": random.choice([-1, 1])
                    }
                }
                await ws.send(json.dumps(npc))
            t2 = time.time()
            print(f"[Stress] Batch {batch+1}: 100 NPCs in {(t2-t1)*1000:.1f}ms")
            await asyncio.sleep(0.1)

        total = time.time() - start
        print(f"[Stress] Total: 5000 tiles + 1000 NPCs in {total:.1f}s")
        print("[Stress] Done")

asyncio.run(main())
