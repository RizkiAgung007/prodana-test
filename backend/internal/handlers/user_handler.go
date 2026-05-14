package handlers

import (
	"backend/internal/config"
	"backend/internal/middleware"
	"backend/internal/models"
	"backend/internal/utils"
	"encoding/json"
	"net/http"
)

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

// CREATE USER
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, `{"message": "Method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Name		string	`json:"name"`
		Email		string	`json:"email"`
		Password	string	`json:"password"`
		RoleID		uint	`json:"role_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"message": "Forma request tidak valid"}`, http.StatusBadRequest)
		return
	}

	// Membuat validasi duplikat email
	var existingUser models.User
	if err := config.DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		http.Error(w, `{"message": "Email sudah digunakan"}`, http.StatusConflict)
		return
	}

	// Melakukan hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		http.Error(w, `{"message": "Gagal memproses password"}`, http.StatusInternalServerError)
		return
	}

	if req.RoleID == 0 {
		req.RoleID = 3
	}

	user := models.User {
		Name	: req.Name,
		Email	: req.Email,
		Password: hashedPassword,
		RoleID	: req.RoleID,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		http.Error(w, `{"message": "Gagal menyimpan user"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"message": "User berhasil ditambahkan.",
	})
}

// GET USER
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		http.Error(w, `{"message": "Method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	var users []models.User
	if err := config.DB.Preload("Role").Find(&users).Error; err != nil {
		http.Error(w, `{"message": "Gagal mengambil data user"}`, http.StatusInternalServerError)
		return
	}

	var response []UserResponse 
	for _, user := range users {
		response = append(response, UserResponse{
			ID	  :	user.ID,
			Name  : user.Name,
			Email : user.Email,
			Role  : user.Role.Name,
		})
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"data": response,
	})
}

// UPDATE USER
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPut {
		http.Error(w, `{"message": "Method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, `{"message": "ID tidak ditemukan"}`, http.StatusBadRequest)
		return
	}

	var req struct {
		Name		string	`json:"name"`
		Email		string	`json:"email"`
		Password	string	`json:"password"`
		RoleID		uint	`json:"role_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"message": "Format request tidak valid"}`, http.StatusBadRequest)
		return
	}

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		http.Error(w, `{"message": "User tidak ditemukan"}`, http.StatusNotFound)
		return
	}

	requestRoleID, ok := r.Context().Value(middleware.RoleIDKey).(uint)
	if !ok {
		http.Error(w, `{"message": "Gagal memuat sesi user"}`, http.StatusUnauthorized)
		return
	}

	if requestRoleID == 2 {
		if user.RoleID != 3 {
			http.Error(w, `{"message": "Akses ditolak, editor hanya bisa mengedit viewer"}`, http.StatusForbidden)
			return
		}

		if req.RoleID != 3 {
			http.Error(w, `{"message": "Akses ditolak, editor tidak bisa mengubah role"}`, http.StatusForbidden)
			return
		}
	}

	user.Name = req.Name
	user.Email = req.Email
	user.RoleID = req.RoleID

	if req.Password != "" {
		hashedPassword, err := utils.HashPassword(req.Password)
		if err != nil {
			http.Error(w, `{"message": "Gagal memproses password"}`, http.StatusInternalServerError)
			return
		}
		
		user.Password = hashedPassword
	}

	if err := config.DB.Save(&user).Error; err != nil {
		http.Error(w, `{"message": "Gagal memperbarui data"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"message": "User berhasil diupdate",
	})
}

// DELETE USER
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodDelete {
		http.Error(w, `{"message": "Method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, `{"message": "ID tidak ditemukan"}`, http.StatusBadRequest)
		return
	}

	if err := config.DB.Delete(&models.User{}, id).Error; err != nil {
		http.Error(w, `{"message": "Gagal menghapus user dari database"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"message": "User berhasil dihapus",
	})
}