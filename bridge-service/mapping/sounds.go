package mapping

type SoundEntry struct {
TRName string
MCName string
}

var Sounds = make(map[int]SoundEntry)

func InitSounds() {
Sounds[0]  = SoundEntry{TRName: "Dig", MCName: "minecraft:block.stone.break"}
Sounds[1]  = SoundEntry{TRName: "PlayerHit", MCName: "minecraft:entity.player.hurt"}
Sounds[2]  = SoundEntry{TRName: "Item1", MCName: "minecraft:entity.item.pickup"}
Sounds[3]  = SoundEntry{TRName: "Item2", MCName: "minecraft:entity.item.pickup"}
Sounds[4]  = SoundEntry{TRName: "Item3", MCName: "minecraft:entity.item.pickup"}
Sounds[5]  = SoundEntry{TRName: "Item4", MCName: "minecraft:entity.item.pickup"}
Sounds[6]  = SoundEntry{TRName: "Item5", MCName: "minecraft:entity.item.pickup"}
Sounds[7]  = SoundEntry{TRName: "Item6", MCName: "minecraft:entity.item.pickup"}
Sounds[8]  = SoundEntry{TRName: "Item7", MCName: "minecraft:entity.item.pickup"}
Sounds[9]  = SoundEntry{TRName: "Item8", MCName: "minecraft:entity.item.pickup"}
Sounds[10] = SoundEntry{TRName: "Item9", MCName: "minecraft:entity.item.pickup"}
Sounds[11] = SoundEntry{TRName: "Item10", MCName: "minecraft:entity.item.pickup"}
Sounds[12] = SoundEntry{TRName: "Item11", MCName: "minecraft:entity.item.pickup"}
Sounds[13] = SoundEntry{TRName: "Item12", MCName: "minecraft:entity.item.pickup"}
Sounds[14] = SoundEntry{TRName: "Item13", MCName: "minecraft:entity.item.pickup"}
Sounds[15] = SoundEntry{TRName: "Item14", MCName: "minecraft:entity.item.pickup"}
Sounds[16] = SoundEntry{TRName: "Item15", MCName: "minecraft:entity.item.pickup"}
Sounds[17] = SoundEntry{TRName: "Item16", MCName: "minecraft:entity.item.pickup"}
Sounds[18] = SoundEntry{TRName: "Item17", MCName: "minecraft:entity.item.pickup"}
Sounds[19] = SoundEntry{TRName: "Item18", MCName: "minecraft:entity.item.pickup"}
Sounds[20] = SoundEntry{TRName: "Item19", MCName: "minecraft:entity.item.pickup"}
Sounds[21] = SoundEntry{TRName: "Item20", MCName: "minecraft:entity.item.pickup"}
Sounds[22] = SoundEntry{TRName: "Item21", MCName: "minecraft:entity.item.pickup"}
Sounds[23] = SoundEntry{TRName: "Item22", MCName: "minecraft:entity.item.pickup"}
Sounds[24] = SoundEntry{TRName: "Item23", MCName: "minecraft:entity.item.pickup"}
Sounds[25] = SoundEntry{TRName: "Item24", MCName: "minecraft:entity.item.pickup"}
Sounds[26] = SoundEntry{TRName: "Item25", MCName: "minecraft:entity.item.pickup"}
Sounds[27] = SoundEntry{TRName: "Item26", MCName: "minecraft:entity.item.pickup"}
Sounds[28] = SoundEntry{TRName: "Item27", MCName: "minecraft:entity.item.pickup"}
Sounds[29] = SoundEntry{TRName: "Item28", MCName: "minecraft:entity.item.pickup"}
Sounds[30] = SoundEntry{TRName: "Item29", MCName: "minecraft:entity.item.pickup"}
Sounds[31] = SoundEntry{TRName: "Item30", MCName: "minecraft:entity.item.pickup"}
Sounds[32] = SoundEntry{TRName: "Item31", MCName: "minecraft:entity.item.pickup"}
Sounds[33] = SoundEntry{TRName: "Item32", MCName: "minecraft:entity.item.pickup"}
Sounds[34] = SoundEntry{TRName: "Item33", MCName: "minecraft:entity.item.pickup"}
Sounds[35] = SoundEntry{TRName: "Item34", MCName: "minecraft:entity.item.pickup"}
Sounds[36] = SoundEntry{TRName: "Item35", MCName: "minecraft:entity.item.pickup"}
Sounds[37] = SoundEntry{TRName: "Item36", MCName: "minecraft:entity.item.pickup"}
Sounds[38] = SoundEntry{TRName: "Item37", MCName: "minecraft:entity.item.pickup"}
Sounds[39] = SoundEntry{TRName: "Item38", MCName: "minecraft:entity.item.pickup"}
Sounds[40] = SoundEntry{TRName: "Item39", MCName: "minecraft:entity.item.pickup"}

Sounds[42] = SoundEntry{TRName: "Killed", MCName: "minecraft:entity.player.death"}
Sounds[43] = SoundEntry{TRName: "Zombie1", MCName: "minecraft:entity.zombie.ambient"}
Sounds[44] = SoundEntry{TRName: "Zombie2", MCName: "minecraft:entity.zombie.hurt"}
Sounds[45] = SoundEntry{TRName: "Zombie3", MCName: "minecraft:entity.zombie.death"}
Sounds[46] = SoundEntry{TRName: "Skeleton1", MCName: "minecraft:entity.skeleton.ambient"}
Sounds[47] = SoundEntry{TRName: "Skeleton2", MCName: "minecraft:entity.skeleton.hurt"}
Sounds[48] = SoundEntry{TRName: "Skeleton3", MCName: "minecraft:entity.skeleton.death"}
Sounds[49] = SoundEntry{TRName: "Slime1", MCName: "minecraft:entity.slime.squish"}
Sounds[50] = SoundEntry{TRName: "Slime2", MCName: "minecraft:entity.slime.death"}

Sounds[51] = SoundEntry{TRName: "DoorOpen", MCName: "minecraft:block.wooden_door.open"}
Sounds[52] = SoundEntry{TRName: "DoorClosed", MCName: "minecraft:block.wooden_door.close"}

Sounds[53] = SoundEntry{TRName: "BowShot", MCName: "minecraft:entity.arrow.shoot"}
Sounds[54] = SoundEntry{TRName: "ArrowHit", MCName: "minecraft:entity.arrow.hit_player"}

Sounds[55] = SoundEntry{TRName: "Splash", MCName: "minecraft:entity.player.splash"}
Sounds[56] = SoundEntry{TRName: "Lava", MCName: "minecraft:block.lava.ambient"}

Sounds[57] = SoundEntry{TRName: "Boss1", MCName: "minecraft:entity.wither.spawn"}
Sounds[58] = SoundEntry{TRName: "Boss2", MCName: "minecraft:entity.wither.death"}
Sounds[59] = SoundEntry{TRName: "Boss3", MCName: "minecraft:entity.ender_dragon.growl"}
Sounds[60] = SoundEntry{TRName: "Boss4", MCName: "minecraft:entity.ender_dragon.death"}

Sounds[61] = SoundEntry{TRName: "Explosion", MCName: "minecraft:entity.generic.explode"}
Sounds[62] = SoundEntry{TRName: "Thunder", MCName: "minecraft:entity.lightning_bolt.thunder"}

Sounds[63] = SoundEntry{TRName: "Rain", MCName: "minecraft:weather.rain"}
Sounds[64] = SoundEntry{TRName: "Wind", MCName: "minecraft:weather.rain"}

Sounds[65] = SoundEntry{TRName: "MenuOpen", MCName: "minecraft:block.chest.open"}
Sounds[66] = SoundEntry{TRName: "MenuClose", MCName: "minecraft:block.chest.close"}
Sounds[67] = SoundEntry{TRName: "MenuTick", MCName: "minecraft:ui.button.click"}

Sounds[68] = SoundEntry{TRName: "PickupItem", MCName: "minecraft:entity.item.pickup"}
Sounds[69] = SoundEntry{TRName: "DropItem", MCName: "minecraft:entity.item.pickup"}

Sounds[70] = SoundEntry{TRName: "Grapple", MCName: "minecraft:entity.fishing_bobber.throw"}
Sounds[71] = SoundEntry{TRName: "Coins", MCName: "minecraft:entity.experience_orb.pickup"}

Sounds[72] = SoundEntry{TRName: "NpcTalk", MCName: "minecraft:entity.villager.ambient"}
Sounds[73] = SoundEntry{TRName: "NpcClose", MCName: "minecraft:entity.villager.no"}

Sounds[74] = SoundEntry{TRName: "Potion", MCName: "minecraft:entity.witch.drink"}
Sounds[75] = SoundEntry{TRName: "Heal", MCName: "minecraft:entity.player.levelup"}
Sounds[76] = SoundEntry{TRName: "Mana", MCName: "minecraft:block.enchantment_table.use"}

Sounds[77] = SoundEntry{TRName: "Summon", MCName: "minecraft:entity.evoker.cast_spell"}
Sounds[78] = SoundEntry{TRName: "WoF", MCName: "minecraft:entity.wither.ambient"}

Sounds[79] = SoundEntry{TRName: "NightStart", MCName: "minecraft:ambient.cave"}
Sounds[80] = SoundEntry{TRName: "DayStart", MCName: "minecraft:entity.player.levelup"}
Sounds[81] = SoundEntry{TRName: "EclipseStart", MCName: "minecraft:entity.wither.spawn"}

Sounds[82] = SoundEntry{TRName: "MechBoss", MCName: "minecraft:entity.iron_golem.repair"}
Sounds[83] = SoundEntry{TRName: "MechSkull", MCName: "minecraft:entity.wither_skeleton.ambient"}
Sounds[84] = SoundEntry{TRName: "MechWorm", MCName: "minecraft:entity.silverfish.ambient"}

Sounds[85] = SoundEntry{TRName: "Critter", MCName: "minecraft:entity.bat.ambient"}
Sounds[86] = SoundEntry{TRName: "Bird", MCName: "minecraft:entity.parrot.ambient"}

Sounds[87] = SoundEntry{TRName: "Fire", MCName: "minecraft:block.fire.ambient"}
Sounds[88] = SoundEntry{TRName: "Water", MCName: "minecraft:block.water.ambient"}

Sounds[89] = SoundEntry{TRName: "SwordSwing", MCName: "minecraft:entity.player.attack.sweep"}
Sounds[90] = SoundEntry{TRName: "SwordHit", MCName: "minecraft:entity.player.attack.crit"}
}

func TRSoundToMC(trID int) string {
if entry, ok := Sounds[trID]; ok {
return entry.MCName
}
return "minecraft:entity.player.hurt"
}
