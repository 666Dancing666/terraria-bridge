import asyncio
import json
import websockets

async def main():
    try:
        async with websockets.connect('ws://localhost:8080/mc') as ws:
            print("[MC] Connected")

            await asyncio.sleep(0.5)

            actions = [
                {"type": "player_move", "payload": {"x": 50.5, "y": 50.0, "z": 0}},
                {"type": "player_action", "payload": {"action": "jump"}},
                {"type": "tile_break", "payload": {"x": 51, "y": 50, "mc_block": "minecraft:dirt"}},
                {"type": "tile_place", "payload": {"x": 52, "y": 52, "mc_block": "minecraft:torch"}},
                {"type": "interact", "payload": {"x": 50, "y": 50, "target": "chest"}},
                {"type": "chat_message", "payload": {"msg": "Hello from Minecraft!"}},
            ]

            for action in actions:
                await ws.send(json.dumps(action))
                print(f"[MC] Sent: {action['type']}")
                await asyncio.sleep(0.2)

            print("[MC] Receiving messages...")
            for _ in range(30):
                try:
                    msg = await asyncio.wait_for(ws.recv(), timeout=1.0)
                    data = json.loads(msg)
                    t = data['type']
                    if t == "world_snapshot_converted":
                        tiles = len(data['payload'].get('tiles', []))
                        print(f"[MC] Snapshot: {tiles} tiles")
                    elif t == "entity_update_converted":
                        etype = data['payload'].get('entity_type', '?')
                        mc_ent = data['payload'].get('mc_entity', '?')
                        print(f"[MC] Entity: {etype} -> {mc_ent}")
                    else:
                        print(f"[MC] Msg: {t}")
                except asyncio.TimeoutError:
                    pass

            print("[MC] Done")

    except Exception as e:
        print(f"[MC] Error: {type(e).__name__}: {e}")

asyncio.run(main())
