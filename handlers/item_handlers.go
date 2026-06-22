package handlers

import (
	"encoding/json"
	"net/http"
	"packages/models"
	"packages/repository"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ItemHandler struct {
	ItemRepo *repository.ItemRepository
}

func NewItemHandler(itemRepo *repository.ItemRepository) *ItemHandler {
	return &ItemHandler{
		ItemRepo: itemRepo,
	}
}

func (h *ItemHandler) CreateItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid request Paylod")
		return
	}

	item.ID = primitive.NewObjectID()
	item.CreatedAt = time.Now()

	err = h.ItemRepo.CreateItem(&item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}
