package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"webServer/data"
	"webServer/internal/recipes/repo"
	"webServer/internal/server"
)

func main() {

	// connect to db
	ctx := context.Background()
	dsn := "host=localhost port=5432 user=postgres password=prikolpronyra dbname=recipe_finder sslmode=disable"
	dbpool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	nativeDB := stdlib.OpenDBFromPool(dbpool)
	sqlxDB := sqlx.NewDb(nativeDB, "pgx")

	r := repo.NewRepo(sqlxDB)

	//nTag := dto.NewTag{
	//	"Хайп",
	//	"☠️",
	//}
	//
	//res, err := r.TagCreate(ctx, &nTag)
	//if err != nil {
	//	log.Fatalf("Error creating tag: %v", err)
	//}
	//fmt.Println(res)

	//ingr := []dto.Ingredient{
	//	{"Колесо", "4 штуки"},
	//	{"Корпус", "1 штука"},
	//	{"Пицца в багажник", "5 коробок"},
	//}
	//fmt.Println(res)
	//
	//steps := []dto.Step{
	//	{1, "Сядь за руль"},
	//	{2, "Кайфуй"},
	//	{3, "Едь"},
	//	{4, "Едь"},
	//	{5, "Едь"},
	//	{6, "Едь"},
	//	{7, "Тормози, куда гонишь"},
	//}
	//
	//tags := []int{
	//	1,
	//}
	//
	//test := dto.NewRecipe{
	//	"Хонда Сивик", tags, 154, 9999,
	//	250000, ingr, steps,
	//	"https://autogeizer.ru/wp-content/uploads/2021/08/honda-civic.jpg",
	//}
	//
	//res2, err := r.RecipeCreate(ctx, &test)
	//if err != nil {
	//	log.Fatalf("Error creating recipe: %v", err)
	//}
	//fmt.Println(res2)

	//ingr = []dto.Ingredient{
	//	{"Страшилка", "1 штука"},
	//	{"Пугалка", "2 столовых ложки"},
	//}
	//
	//steps = []dto.Step{
	//	{1, "Страшно"},
	//}
	//
	//tags = []int{
	//	1, 2, 4,
	//}
	//
	//test = dto.NewRecipe{
	//	"Ёжиг", tags, 666, 111,
	//	1500, ingr, steps,
	//	"https://cs10.pikabu.ru/post_img/big/2019/11/06/8/1573044991158234281.jpg",
	//}
	//
	//res3, err := r.RecipeCreate(ctx, &test)
	//if err != nil {
	//	log.Fatalf("Error creating recipe: %v", err)
	//}
	//fmt.Println(res3)

	//res3, err := r.TagsAll(ctx)
	//if err != nil {
	//	log.Fatalf("Error fetching tags: %v", err)
	//}
	//
	//for _, tag := range res3 {
	//	fmt.Println(tag)
	//}
	//
	//res4, err := r.RecipesAll(ctx)
	//if err != nil {
	//	log.Fatalf("Error fetching recipes: %v", err)
	//}
	//
	//for _, recipe := range res4 {
	//	fmt.Println(recipe)
	//}

	res5, err := r.RecipeByName(ctx, "fndfmnghm")
	if err != nil {
		log.Fatalf("Error fetching recipe: %v", err)
	}

	fmt.Println(res5)

	res6, err := r.RecipesByBudget(ctx, 175)
	if err != nil {
		log.Fatalf("Error fetching recipes: %v", err)
	}

	for _, recipe := range res6 {
		fmt.Println(recipe)
	}

	res7, err := r.RecipesByTags(ctx, 1, 2, 4)
	if err != nil {
		log.Fatalf("Error fetching recipes: %v", err)
	}

	for _, recipe := range res7 {
		fmt.Println(recipe)
	}

	server.Server(ctx, r)

	//// execute repo
	//
	//router := gin.Default()
	//
	//sT := data.NewTagStore()
	//tagsHandler := NewTagHandler(sT)
	//
	//sS := data.NewSmallStore()
	//smallHandler := NewSmallHandler(sS)
	//
	//sF := data.NewFullStore()
	//fullHandler := NewFullHandler(sF)
	//
	//router.GET("/tags", tagsHandler.GetTags)
	//router.GET("/preview", smallHandler.GetSmall)
	//router.GET("/details", fullHandler.GetFull)
	//
	//router.Run(":8080")

}

type FullHandler struct {
	store fullStore
}

type fullStore interface {
	Get() (data.RecipeDetails, error)
}

func (h FullHandler) GetFull(c *gin.Context) {
	sm, err := h.store.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, sm)
}

func NewFullHandler(s fullStore) *FullHandler {
	return &FullHandler{store: s}
}

type SmallHandler struct {
	store smallStore
}

type smallStore interface {
	Get() ([]data.RecipePreview, error)
}

func (h SmallHandler) GetSmall(c *gin.Context) {
	sm, err := h.store.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, sm)
}

func NewSmallHandler(s smallStore) *SmallHandler {
	return &SmallHandler{store: s}
}

type TagsHandler struct {
	store tagStore
}

type tagStore interface {
	Get() ([]data.Tag, error)
}

func (h TagsHandler) GetTags(c *gin.Context) {
	t, err := h.store.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, t)
}

func NewTagHandler(s tagStore) *TagsHandler {
	return &TagsHandler{
		store: s,
	}
}
