package handlers

import (
	"backend/internal/config"
	"backend/internal/middleware"
	"backend/internal/models"
	"bytes"
	"context"
	"encoding/json"
	"fmt" 
	"net/http"
	"net/http/httptest"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupUserTestDB() {
	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	config.DB = db
	db.AutoMigrate(&models.Role{}, &models.User{})
	
	db.Exec("DELETE FROM users")
	db.Exec("DELETE FROM roles")

	db.Create(&models.Role{ID: 1, Name: "Admin"})
	db.Create(&models.Role{ID: 2, Name: "Editor"})
	db.Create(&models.Role{ID: 3, Name: "Viewer"})
}

// TestCreateUser_Success pembuatan user baru oleh Admin
func TestCreateUser_Success(t *testing.T) {
	setupUserTestDB()

	reqBody := []byte(`{"name": "Ahmad", "email": "ahmad12@mail.com", "password": "password123", "role_id": 3}`)
	req, _ := http.NewRequest("POST", "/api/users", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	ctx := context.WithValue(req.Context(), middleware.RoleIDKey, uint(1))
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateUser)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Harusnya status 201, dapat %v", status)
	}

	var user models.User
	config.DB.Where("email = ?", "ahmad12@mail.com").First(&user)
	if user.Name != "Ahmad" {
		t.Errorf("Nama user tidak sesuai, dapat %v", user.Name)
	}
}

// TestGetUsers mengambil daftar user
func TestGetUsers(t *testing.T) {
	setupUserTestDB()

	config.DB.Create(&models.User{Name: "Agus", Email: "agus@mail.com", RoleID: 3})

	req, _ := http.NewRequest("GET", "/api/users", nil)
	rr := httptest.NewRecorder()
	
	handler := http.HandlerFunc(GetUsers)
	handler.ServeHTTP(rr, req)

	var response map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &response)

	if response["status"] != "success" {
		t.Errorf("Status harusnya success")
	}

	data := response["data"].([]interface{})
	if len(data) == 0 {
		t.Errorf("Data user harusnya tidak kosong")
	}
}

// TestUpdateUser_Forbidden jika Editor mencoba mengedit Admin (Harus Gagal)
func TestUpdateUser_Forbidden(t *testing.T) {
	setupUserTestDB()

	admin := models.User{Name: "Real Admin", Email: "admin@mail.com", RoleID: 1}
	config.DB.Create(&admin)

	reqBody := []byte(`{"name": "Admin Diubah", "email": "admin@mail.com", "role_id": 1}`)
	
	url := fmt.Sprintf("/api/users?id=%d", admin.ID)
	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(reqBody))
	
	ctx := context.WithValue(req.Context(), middleware.RoleIDKey, uint(2))
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateUser)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusForbidden {
		t.Errorf("Harusnya status 403 (Forbidden), tapi dapat %v", status)
	}
}

// TestDeleteUser penghapusan user
func TestDeleteUser(t *testing.T) {
	setupUserTestDB()
	
	userToDelete := models.User{Name: "User Dihapus", Email: "hapus@mail.com", RoleID: 3}
	config.DB.Create(&userToDelete)

	url := fmt.Sprintf("/api/users?id=%d", userToDelete.ID)
	req, _ := http.NewRequest("DELETE", url, nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(DeleteUser)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Gagal menghapus user, status: %v", status)
	}

	var user models.User
	result := config.DB.First(&user, userToDelete.ID)
	if result.Error == nil {
		t.Errorf("User harusnya sudah tidak ada di database")
	}
}