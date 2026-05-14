package handlers

import (
	"backend/internal/config"
	"backend/internal/models"
	"encoding/json"
	"net/http"
)

// CREATE PRODUCT
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, `{"message": "Method not allowed"}`, http.StatusInternalServerError)
		return
	}

	var products models.Product
	if err := json.NewDecoder(r.Body).Decode(&products); err != nil {
		http.Error(w, `{"message": "Format request tidak valid"}`, http.StatusBadRequest)
		return
	}

	if err := config.DB.Create(&products).Error; err != nil {
		http.Error(w, `{"message": "Gagal menyimpan produk"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"message": "Produk berhasil ditambahkan",
		"data": products,
	})
}

// GET PRODUCT
func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		http.Error(w, `{"message": "Method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	var products []models.Product
	if err := config.DB.Find(&products).Error; err != nil {
		http.Error(w, `{"message": "Gagal mengambil data produk"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"data": products,
	})
}

// UPDATE PRODUCT
func UpdateProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPut {
		http.Error(w, `{"message": "Method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, `{"message": "ID tidka ditemukan"}`, http.StatusBadRequest)
		return
	}

	var req struct {
		Name		string	`json:"name"`
		Price		float64	`json:"price"`
		Stock		int		`json:"stock"`
		Description string	`json:"description"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"message": "Format request tidak valid"}`, http.StatusBadRequest)
		return
	}

	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		http.Error(w, `{"message": "Produk tidak ditemukan"}`, http.StatusNotFound)
		return
	}

	product.Name = req.Name
	product.Price = req.Price
	
	if err := config.DB.Save(&product).Error; err != nil {
		http.Error(w, `{"message": "Gagal memperbarui produk"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"message": "Produk berhasil diperbarui",
	})
}

// DELETE PRODUCT
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
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

	if err := config.DB.Delete(&models.Product{}, id).Error; err != nil {
		http.Error(w, `{"message": "Gagal menhapus produk"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"message": "Produk berhasil dihapus",
	})
}