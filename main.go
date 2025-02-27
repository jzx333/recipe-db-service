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

	res, err := r.TagsAll(ctx)
	if err != nil {
		log.Fatalf("Error fetching tags: %v", err)
	}

	for _, tag := range res {
		fmt.Println("%+v\n", *tag)
	}
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
