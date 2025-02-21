package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"webServer/data"
)

func main() {
	router := gin.Default()

	store := data.NewTagStore()
	tagsHandler := NewTagHandler(store)

	router.GET("/tags", tagsHandler.GetTags)

	router.Run("localhost:8080")

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
