package handlers

import (
	"encoding/json"
	"net/http"
	"packages/models"
	"packages/repository"
	"packages/utils"
	"time"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
	UserRepo *repository.UserRepository
}

func NewUserHandler(userRepo *repository.UserRepository) *UserHandler {
	return &UserHandler{
		UserRepo: userRepo,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid request Paylod")
		return
	}

	user.ID = primitive.NewObjectID()
	user.Password = "ItemMigration@123"
	user.IsActive = true
	user.IsDeleted = false
	user.CreatedAt = time.Now()

	err = h.UserRepo.CreateUser(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// 6a1d1ab31d984f941362aca2
func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {

	// URL nundi id teesuko
	id := chi.URLParam(r, "id")
	// string -> ObjectID convert cheyyi
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid requestID")
		return
	}
	// repository call cheyyi
	user, err := h.UserRepo.GetUserByID(objectID)

	// error handle cheyyi
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("User Didn't existed")
		return
	}
	// success response pampu
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {

	allUsers, err := h.UserRepo.GetAllUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("There are no  Users to retrieve")
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(allUsers)
}

func (h *UserHandler) DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("There are no  Users to retrieve")
		return
	}
	err = h.UserRepo.DeleteUserByID(objectID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Failed to delete user")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("User Deleted Successfully")
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Request")
		return
	}
	user, err := h.UserRepo.GetUserByEmail(req.Email)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("User Doesn't Exists")
		return
	}
	if user.Password != req.Password {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Invalid Password")
		return
	}
	token, err := utils.GenerateToken(user.ID.Hex())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Failed To Generate Token")
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User Login Successfull",
		"token":   token,
	})
}

func (h *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Logout successful",
	})
}

func (h *UserHandler) ResetPassword(w http.ResponseWriter, r *http.Request) {
	var req models.ResetPassword
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Request")
		return
	}
	id := chi.URLParam(r, "id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid requestID")
		return
	}
	user, err := h.UserRepo.GetUserByID(objectID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid requestID")
		return
	}

	if user.Password != req.NewPassword && req.NewPassword != req.OldPassword {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Old Password Incorrect or Using Same Password")
		return
	}

	err = h.UserRepo.UpdatePassword(objectID, req.NewPassword)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Failed To Update Password")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Password Updated Successfull",
	})

}
