package mapping

type TileEntry struct {
TRName string `json:"tr_name"`
MCName string `json:"mc_name"`
}

type ItemEntry struct {
TRName string `json:"tr_name"`
MCName string `json:"mc_name"`
}

type EntityEntry struct {
TRName string `json:"tr_name"`
MCName string `json:"mc_name"`
}

var (
Tiles    = make(map[int]TileEntry)
Items    = make(map[int]ItemEntry)
Entities = make(map[int]EntityEntry)
Walls    = make(map[int]TileEntry)
)

func TRTileToMC(trID int) string {
if entry, ok := Tiles[trID]; ok {
return entry.MCName
}
return "terrariabridge:unknown_block"
}

func MCTileToTR(mcName string) int {
for id, entry := range Tiles {
if entry.MCName == mcName {
return id
}
}
return 0
}

func TRWallToMC(trID int) string {
if entry, ok := Walls[trID]; ok {
return entry.MCName
}
return "terrariabridge:unknown_wall"
}

func TRItemToMC(trID int) string {
if entry, ok := Items[trID]; ok {
return entry.MCName
}
return "minecraft:stick"
}

func MCItemToTR(mcName string) int {
for id, entry := range Items {
if entry.MCName == mcName {
return id
}
}
return 0
}

func TREntityToMC(trID int) string {
if entry, ok := Entities[trID]; ok {
return entry.MCName
}
return "terrariabridge:unknown_entity"
}
