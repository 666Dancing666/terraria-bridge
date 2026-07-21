package mapping

type HookEntry struct {
TRName    string
Range     float64
Speed     float64
Hooks     int
Latching  bool
}

var Hooks = make(map[int]HookEntry)

func InitHooks() {
Hooks[52]  = HookEntry{TRName: "Grappling Hook", Range: 18, Speed: 11, Hooks: 1, Latching: false}
Hooks[53]  = HookEntry{TRName: "Sapphire Hook", Range: 16, Speed: 10, Hooks: 1, Latching: false}
Hooks[54]  = HookEntry{TRName: "Emerald Hook", Range: 18, Speed: 11, Hooks: 1, Latching: false}
Hooks[55]  = HookEntry{TRName: "Ruby Hook", Range: 20, Speed: 12, Hooks: 1, Latching: false}
Hooks[56]  = HookEntry{TRName: "Diamond Hook", Range: 22, Speed: 13, Hooks: 1, Latching: false}
Hooks[57]  = HookEntry{TRName: "Amethyst Hook", Range: 15, Speed: 10, Hooks: 1, Latching: false}
Hooks[58]  = HookEntry{TRName: "Topaz Hook", Range: 14, Speed: 9, Hooks: 1, Latching: false}
Hooks[59]  = HookEntry{TRName: "Dual Hook", Range: 21, Speed: 13, Hooks: 2, Latching: false}
Hooks[60]  = HookEntry{TRName: "Spooky Hook", Range: 25, Speed: 14, Hooks: 1, Latching: false}
Hooks[61]  = HookEntry{TRName: "Christmas Hook", Range: 22, Speed: 12, Hooks: 1, Latching: false}
Hooks[62]  = HookEntry{TRName: "Bat Hook", Range: 25, Speed: 14, Hooks: 1, Latching: false}
Hooks[63]  = HookEntry{TRName: "Skeletron Hand", Range: 20, Speed: 12, Hooks: 1, Latching: false}
Hooks[64]  = HookEntry{TRName: "Fish Hook", Range: 18, Speed: 12, Hooks: 1, Latching: false}
Hooks[65]  = HookEntry{TRName: "Web Hook", Range: 18, Speed: 10, Hooks: 8, Latching: true}
Hooks[66]  = HookEntry{TRName: "Slime Hook", Range: 16, Speed: 10, Hooks: 1, Latching: false}
Hooks[67]  = HookEntry{TRName: "Lunar Hook", Range: 25, Speed: 15, Hooks: 4, Latching: false}
Hooks[68]  = HookEntry{TRName: "Static Hook", Range: 28, Speed: 14, Hooks: 1, Latching: true}
Hooks[69]  = HookEntry{TRName: "Anti-Gravity Hook", Range: 25, Speed: 13, Hooks: 1, Latching: false}
Hooks[70]  = HookEntry{TRName: "Thorn Hook", Range: 22, Speed: 13, Hooks: 1, Latching: false}
Hooks[71]  = HookEntry{TRName: "Tendon Hook", Range: 20, Speed: 12, Hooks: 1, Latching: false}
Hooks[72]  = HookEntry{TRName: "Illuminant Hook", Range: 22, Speed: 13, Hooks: 1, Latching: false}
Hooks[73]  = HookEntry{TRName: "Worm Hook", Range: 18, Speed: 12, Hooks: 1, Latching: false}
Hooks[74]  = HookEntry{TRName: "Silk Hook", Range: 18, Speed: 11, Hooks: 1, Latching: false}
Hooks[75]  = HookEntry{TRName: "Ivy Whip", Range: 20, Speed: 12, Hooks: 3, Latching: false}
}

func GetHook(itemID int) *HookEntry {
if hook, ok := Hooks[itemID]; ok {
return &hook
}
return nil
}
