import asyncio
import json
import websockets
import random

async def main():
    try:
        async with websockets.connect('ws://localhost:8080/tshock') as ws:
            print("[TShock] Connected")

            snapshot = {
                "type": "world_snapshot",
                "payload": {
                    "center_x": 50, "center_y": 50,
                    "radius_x": 10, "radius_y": 10,
                    "world_width": 100, "world_height": 100,
                    "tiles": [
                        {"x": 50, "y": 50, "tile_type": 1, "wall_type": 1},
                        {"x": 51, "y": 50, "tile_type": 2},
                        {"x": 52, "y": 50, "tile_type": 5, "wall_type": 4},
                        {"x": 50, "y": 51, "tile_type": 1},
                        {"x": 51, "y": 51, "tile_type": 38, "wall_type": 5},
                        {"x": 52, "y": 51, "tile_type": 30},
                        {"x": 50, "y": 52, "tile_type": 3},
                        {"x": 51, "y": 52, "tile_type": 22},
                        {"x": 52, "y": 52, "tile_type": 54},
                    ]
                }
            }
            await ws.send(json.dumps(snapshot))
            print("[TShock] Sent world snapshot")

            await asyncio.sleep(0.5)

            for i in range(5):
                entity = {
                    "type": "entity_update",
                    "payload": {
                        "entity_type": "npc",
                        "id": i,
                        "net_id": random.choice([1, 3, 6, 26, 43]),
                        "x": 50 * 16 + random.randint(-80, 80),
                        "y": 50 * 16 + random.randint(-80, 80),
                        "health": 100,
                        "max_health": 100,
                        "direction": random.choice([-1, 1])
                    }
                }
                await ws.send(json.dumps(entity))
            print("[TShock] Sent 5 NPCs")

            await asyncio.sleep(0.5)

            player = {
                "type": "entity_update",
                "payload": {
                    "entity_type": "player",
                    "id": 0,
                    "name": "TestPlayer",
                    "x": 50 * 16,
                    "y": 50 * 16,
                    "health": 200,
                    "max_health": 200,
                    "direction": 1
                }
            }
            await ws.send(json.dumps(player))
            print("[TShock] Sent player")

            for i in range(3):
                item = {
                    "type": "entity_update",
                    "payload": {
                        "entity_type": "item",
                        "id": 100 + i,
                        "net_id": random.choice([22, 50, 75, 73]),
                        "x": 50 * 16 + i * 30,
                        "y": 50 * 16 + 10,
                        "stack": random.randint(1, 99)
                    }
                }
                await ws.send(json.dumps(item))
            print("[TShock] Sent 3 items")

            await asyncio.sleep(2)
            print("[TShock] Done")

    except Exception as e:
        print(f"[TShock] Error: {type(e).__name__}: {e}")

asyncio.run(main())
