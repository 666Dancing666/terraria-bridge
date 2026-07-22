using System.Collections.Generic;
using Terraria;
using TerrariaBridge.Network;

namespace TerrariaBridge.Handlers
{
    public class EntityHandler
    {
        private BridgeClient _client;

        public EntityHandler(BridgeClient client)
        {
            _client = client;
        }

        public BridgeMessage CreatePlayerUpdate(Player player)
        {
            return new BridgeMessage
            {
                Type = MessageType.EntityUpdate,
                Payload = new Dictionary<string, object>
                {
                    { "entity_type", "player" },
                    { "id", player.whoAmI },
                    { "name", player.name },
                    { "x", player.position.X },
                    { "y", player.position.Y },
                    { "velocity_x", player.velocity.X },
                    { "velocity_y", player.velocity.Y },
                    { "health", player.statLife },
                    { "max_health", player.statLifeMax2 },
                    { "direction", player.direction }
                }
            };
        }

        public BridgeMessage CreateNPCUpdate(NPC npc)
        {
            return new BridgeMessage
            {
                Type = MessageType.EntityUpdate,
                Payload = new Dictionary<string, object>
                {
                    { "entity_type", "npc" },
                    { "id", npc.whoAmI },
                    { "net_id", npc.netID },
                    { "x", npc.position.X },
                    { "y", npc.position.Y },
                    { "velocity_x", npc.velocity.X },
                    { "velocity_y", npc.velocity.Y },
                    { "health", npc.life },
                    { "max_health", npc.lifeMax },
                    { "direction", npc.direction }
                }
            };
        }

        public BridgeMessage CreateItemUpdate(Item item, int index)
        {
            return new BridgeMessage
            {
                Type = MessageType.EntityUpdate,
                Payload = new Dictionary<string, object>
                {
                    { "entity_type", "item" },
                    { "id", index },
                    { "net_id", item.type },
                    { "stack", item.stack },
                    { "prefix", item.prefix },
                    { "x", 0f },
                    { "y", 0f },
                }
            };
        }

        public BridgeMessage CreateProjectileUpdate(Projectile proj)
        {
            return new BridgeMessage
            {
                Type = MessageType.EntityUpdate,
                Payload = new Dictionary<string, object>
                {
                    { "entity_type", "projectile" },
                    { "id", proj.whoAmI },
                    { "net_id", proj.type },
                    { "x", proj.position.X },
                    { "y", proj.position.Y },
                    { "velocity_x", proj.velocity.X },
                    { "velocity_y", proj.velocity.Y },
                    { "owner", proj.owner }
                }
            };
        }

        public BridgeMessage CreateEntityRemove(string entityType, int id)
        {
            return new BridgeMessage
            {
                Type = MessageType.EntityRemove,
                Payload = new Dictionary<string, object>
                {
                    { "entity_type", entityType },
                    { "id", id }
                }
            };
        }

        public BridgeMessage CreatePlayerJoin(Player player)
        {
            return new BridgeMessage
            {
                Type = MessageType.PlayerJoin,
                Payload = new Dictionary<string, object>
                {
                    { "id", player.whoAmI },
                    { "name", player.name },
                    { "x", player.position.X },
                    { "y", player.position.Y }
                }
            };
        }

        public BridgeMessage CreatePlayerLeave(int id)
        {
            return new BridgeMessage
            {
                Type = MessageType.PlayerLeave,
                Payload = new Dictionary<string, object>
                {
                    { "id", id }
                }
            };
        }
    }
}
