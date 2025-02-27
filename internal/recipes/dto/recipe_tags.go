package dto

type RecipeTags struct {
	RecipeId int `db:"recipe_id"`
	TagId    int `db:"tag_id"`
}
