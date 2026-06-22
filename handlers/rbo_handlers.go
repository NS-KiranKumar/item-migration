package handlers

import (
	"encoding/json"
	"net/http"
	"packages/models"
	"packages/repository"
	"time"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RboHandler struct {
	RboRepo *repository.RboRepository
}

func NewRboHandler(rboRepo *repository.RboRepository) *RboHandler {
	return &RboHandler{
		RboRepo: rboRepo,
	}
}

func (h *RboHandler) CreateRbo(w http.ResponseWriter, r *http.Request) {
	var rbo models.RBO
	err := json.NewDecoder(r.Body).Decode(&rbo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid request Paylod")
		return
	}
	rbo.ID = primitive.NewObjectID()
	rbo.CreatedAt = time.Now()
	err = h.RboRepo.CreateRbo(&rbo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rbo)
}

func (h *RboHandler) GetRbobyID(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid requestID")
		return
	}

	rbo, err := h.RboRepo.GetRbobyID(objectID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Failed to get RBO")
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rbo)

}

func (h *RboHandler) DeleteRbobyID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid requestID")
		return
	}
	err = h.RboRepo.DeleteRbobyID(objectID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Failed to delete RBO")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("RBO Deleted Successfully")
}

func (h *RboHandler) UpdateRbobyID(w http.ResponseWriter, r *http.Request) {
	var rbo models.RBO
	err := json.NewDecoder(r.Body).Decode(&rbo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Request Body")
		return
	}
	id := chi.URLParam(r, "id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid requestID")
		return
	}

	rbo.UpdatedAt = time.Now()

	err = h.RboRepo.UpdateRbobyID(objectID, &rbo)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Failed to update RBO")
		return
	}
	updatedRBO, err := h.RboRepo.GetRbobyID(objectID)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedRBO)
}
func (h *RboHandler) GetAllRbos(w http.ResponseWriter, r *http.Request) {
	AllRbos, err := h.RboRepo.GetallRbos()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid requestID")
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(AllRbos)
}
