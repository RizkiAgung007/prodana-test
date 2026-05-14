package handlers

import (
	"backend/internal/config"
	"backend/internal/models"
	"backend/internal/utils"
	"encoding/json"
	"net/http"
)

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleID   uint   `json:"role_id"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// REGISTER
func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// METHOD POST
	if r.Method != http.MethodPost {
		http.Error(w, `{"message": "Method not allowed}`, http.StatusMethodNotAllowed)
		return
	}

	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"message": "Format request tidak valid"}`, http.StatusBadRequest)
		return
	}

	// Melakukan pengecekan apakah email sudah terdaftar atau belum
	var existingUser models.User
	if err := config.DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		http.Error(w, `{"message": "Email sudah terdaftar"}`, http.StatusConflict)
	}

	// Melakukan hash passwrod
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		http.Error(w, `{"message": "Gagal melakukan hash password"}`, http.StatusInternalServerError)
		return
	}

	// Data default dengan role id 3 (VIEWER)
	roleID := req.RoleID
	if roleID == 0 {
		roleID = 3
	}

	user := models.User {
		Name:	req.Name,
		Email:	req.Email,
		Password: hashedPassword,
		RoleID: roleID,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		http.Error(w, `{"message": "GGagal menyimpan user"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"message": "User berhasil dibuat",
	})
}

// LOGIN
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, `{"message": "Method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"message": "Format request tidak valid"}`, http.StatusBadRequest)
		return
	}

	// Melakukan pencarian berdasarkan email
	var user models.User
	if err := config.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		http.Error(w, `{"message": "Email atau password salah"}`, http.StatusUnauthorized)
		return
	}

	// Mengecek passwrod
	if match := utils.CheckPasswordHash(req.Password, user.Password); !match {
		http.Error(w, `{"message": "Email atau password salah"}`, http.StatusUnauthorized)
		return
	}

	// Membuat tuoken JWT
	token, err := utils.GenerateJwt(user.ID, user.RoleID)
	if err != nil {
		http.Error(w, `{"message": "Gagal membuat token"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"message": "Login Berhasil",
		"data": map[string]interface{}{
			"token": token,
			"user": map[string]interface{}{
				"id"	 : user.ID,
				"name"	 : user.Name,
				"email"	 : user.Email,
				"role_id": user.RoleID,
			},
		},
	})
}