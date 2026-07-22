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
                    { "mana", player.statMana },
                    { "max_mana", player.statManaMax2 },
                    { "direction", player.direction },
                    { "selected_item", player.selectedItem },
                    { "team", player.team }
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
                    { "direction", npc.direction },
                    { "ai_0", npc.ai[0] },
                    { "ai_1", npc.ai[1] },
                    { "ai_2", npc.ai[2] },
                    { "ai_3", npc.ai[3] },
                    { "target", npc.target }
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
                    { "x", item.Center.X },
                    { "y", item.Center.Y },
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
                    { "net_id", proj.netID },
                    { "type", proj.type },
                    { "x", proj.position.X },
                    { "y", proj.position.Y },
                    { "velocity_x", proj.velocity.X },
                    { "velocity_y", proj.velocity.Y },
                    { "owner", proj.owner },
                    { "damage", proj.damage },
                    { "ai_0", proj.ai[0] },
                    { "ai_1", proj.ai[1] }
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
