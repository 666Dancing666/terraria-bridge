package mapping

type RecipeEntry struct {
ResultItem  int
ResultCount int
Station     string
Ingredients map[int]int
}

var Recipes = make(map[string]RecipeEntry)

func InitRecipes() {
Recipes["torch"] = RecipeEntry{
ResultItem: 50, ResultCount: 3,
Station: "workbench",
Ingredients: map[int]int{95: 1, 8: 1},
}

Recipes["wood_platform"] = RecipeEntry{
ResultItem: 9, ResultCount: 2,
Station: "workbench",
Ingredients: map[int]int{8: 1},
}

Recipes["wood_door"] = RecipeEntry{
ResultItem: 26, ResultCount: 1,
Station: "workbench",
Ingredients: map[int]int{8: 6},
}

Recipes["wood_chest"] = RecipeEntry{
ResultItem: 48, ResultCount: 1,
Station: "workbench",
Ingredients: map[int]int{8: 8, 22: 2},
}

Recipes["furnace"] = RecipeEntry{
ResultItem: 35, ResultCount: 1,
Station: "workbench",
Ingredients: map[int]int{1: 20, 8: 4, 50: 3},
}

Recipes["iron_anvil"] = RecipeEntry{
ResultItem: 34, ResultCount: 1,
Station: "workbench",
Ingredients: map[int]int{22: 5},
}

Recipes["iron_pickaxe"] = RecipeEntry{
ResultItem: 1, ResultCount: 1,
Station: "anvil",
Ingredients: map[int]int{22: 12, 8: 4},
}

Recipes["iron_sword"] = RecipeEntry{
ResultItem: 4, ResultCount: 1,
Station: "anvil",
Ingredients: map[int]int{22: 8, 8: 4},
}

Recipes["wood_bow"] = RecipeEntry{
ResultItem: 39, ResultCount: 1,
Station: "workbench",
Ingredients: map[int]int{8: 10},
}

Recipes["wood_arrow"] = RecipeEntry{
ResultItem: 40, ResultCount: 25,
Station: "workbench",
Ingredients: map[int]int{8: 1, 1: 1},
}

Recipes["flaming_arrow"] = RecipeEntry{
ResultItem: 41, ResultCount: 10,
Station: "workbench",
Ingredients: map[int]int{40: 10, 50: 1},
}

Recipes["lesser_healing"] = RecipeEntry{
ResultItem: 100, ResultCount: 1,
Station: "bottle",
Ingredients: map[int]int{31: 1, 95: 2},
}

Recipes["healing_potion"] = RecipeEntry{
ResultItem: 101, ResultCount: 1,
Station: "bottle",
Ingredients: map[int]int{100: 2, 109: 1},
}

Recipes["bomb"] = RecipeEntry{
ResultItem: 86, ResultCount: 1,
Station: "workbench",
Ingredients: map[int]int{22: 3, 95: 5},
}

Recipes["stone_wall"] = RecipeEntry{
ResultItem: 27, ResultCount: 4,
Station: "workbench",
Ingredients: map[int]int{3: 1},
}

Recipes["wood_wall"] = RecipeEntry{
ResultItem: 10, ResultCount: 4,
Station: "workbench",
Ingredients: map[int]int{8: 1},
}

Recipes["chain"] = RecipeEntry{
ResultItem: 22, ResultCount: 10,
Station: "anvil",
Ingredients: map[int]int{22: 1},
}

Recipes["grappling_hook"] = RecipeEntry{
ResultItem: 52, ResultCount: 1,
Station: "anvil",
Ingredients: map[int]int{22: 3, 85: 1},
}

Recipes["chest"] = RecipeEntry{
ResultItem: 48, ResultCount: 1,
Station: "workbench",
Ingredients: map[int]int{8: 8, 22: 2},
}

Recipes["bed"] = RecipeEntry{
ResultItem: 206, ResultCount: 1,
Station: "workbench",
Ingredients: map[int]int{8: 15, 85: 5},
}
}

func CheckRecipe(mcItem string) *RecipeEntry {
if recipe, ok := Recipes[mcItem]; ok {
return &recipe
}
return nil
}
