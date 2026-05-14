package handlers

import (
	"backend/internal/config"
	"backend/internal/models"
	"bytes"
	"encoding/json"
	"fmt" 
	"net/http"
	"net/http/httptest"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("Gagal menghubungkan database test")
	}

	config.DB = db
	db.AutoMigrate(&models.Product{})
	
	db.Exec("DELETE FROM products")
}

// TEST GET /api/products
func TestGetProducts(t *testing.T) {
	setupTestDB()

	dummyProduct := models.Product{Name: "Kursi Plastik", Price: 1500000}
	config.DB.Create(&dummyProduct)

	req, err := http.NewRequest("GET", "/api/products", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(GetProducts)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler mengembalikan status code salah: dapat %v, seharusnya %v", status, http.StatusOK)
	}

	var response map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &response)

	if response["status"] != "success" {
		t.Errorf("Ekspektasi status 'success', tapi dapat %v", response["status"])
	}
}

// TEST POST /api/products
func TestCreateProduct(t *testing.T) {
	setupTestDB()

	reqBody := []byte(`{"name": "Monitor 24 Inch", "price": 2000000}`)

	req, err := http.NewRequest("POST", "/api/products", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateProduct)
	
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Dapat status code %v, seharusnya %v", status, http.StatusCreated)
	}

	var count int64
	config.DB.Model(&models.Product{}).Count(&count)
	if count != 1 {
		t.Errorf("Ekspektasi 1 produk di database, tapi ada %v", count)
	}

	var savedProduct models.Product
	config.DB.First(&savedProduct)
	if savedProduct.Name != "Monitor 24 Inch" {
		t.Errorf("Ekspektasi nama 'Monitor 24 Inch', tapi di DB tersimpan '%v'", savedProduct.Name)
	}
}

// TEST PUT /api/products
func TestUpdateProduct(t *testing.T) {
	setupTestDB()

	initialProduct := models.Product{Name: "Mouse Biasa", Price: 50000}
	config.DB.Create(&initialProduct)

	reqBody := []byte(`{"name": "Mouse Gaming", "price": 300000}`)

	url := fmt.Sprintf("/api/products?id=%d", initialProduct.ID)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	
	handler := http.HandlerFunc(UpdateProducts)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Dapat status code %v, seharusnya %v", status, http.StatusOK)
	}

	var updatedProduct models.Product
	config.DB.First(&updatedProduct, initialProduct.ID) 
	if updatedProduct.Name != "Mouse Gaming" {
		t.Errorf("Ekspektasi nama berubah jadi 'Mouse Gaming', tapi di DB '%v'", updatedProduct.Name)
	}
	
	if updatedProduct.Price != 300000 {
		t.Errorf("Ekspektasi harga berubah jadi 300000, tapi di DB '%v'", updatedProduct.Price)
	}
}

// TEST DELETE /api/products
func TestDeleteProduct(t *testing.T) {
	setupTestDB()

	productToDelete := models.Product{Name: "Keyboard Rusak", Price: 10000}
	config.DB.Create(&productToDelete)

	var initialCount int64
	config.DB.Model(&models.Product{}).Count(&initialCount)
	if initialCount != 1 {
		t.Fatalf("Gagal melakukan setup data, jumlah produk di DB: %v", initialCount)
	}

	url := fmt.Sprintf("/api/products?id=%d", productToDelete.ID)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteProduct)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Dapat status code %v, seharusnya %v", status, http.StatusOK)
	}

	var finalCount int64
	config.DB.Model(&models.Product{}).Count(&finalCount)
	if finalCount != 0 {
		t.Errorf("Ekspektasi produk di DB adalah 0 (terhapus), tapi masih ada %v", finalCount)
	}
}