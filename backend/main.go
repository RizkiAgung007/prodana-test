package main

import (
	"backend/internal/config"
	"backend/internal/handlers"
	"backend/internal/middleware"
	"backend/internal/models"
	"log"
	"net/http"
)

func seedRoles() {
	roles := []string{"Admin", "Editor", "Viewer"}
	
	for _, roleName := range roles {
		var role models.Role

		err := config.DB.Where("name = ?", roleName).First(&role).Error
		if err != nil {
			config.DB.Create(&models.Role{Name: roleName})
			log.Printf("Role '%s' berhasil dibuat\n", roleName)
		}
	}
}

func main() {
	config.ConnectDB()
	seedRoles()

	mux := http.NewServeMux()

	// AUTH
	mux.HandleFunc("/api/register", handlers.Register)
	mux.HandleFunc("/api/login", handlers.Login)

	// CRUD USER
	mux.HandleFunc("/api/users", middleware.AuthMiddleware(middleware.RoleMiddleware(1, 2, 3)(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			middleware.RoleMiddleware(1)(handlers.CreateUser)(w, r)
		} else if r.Method == http.MethodPut {
			middleware.RoleMiddleware(1, 2)(handlers.UpdateUser)(w, r)
		} else if r.Method == http.MethodGet {
			middleware.RoleMiddleware(1, 2, 3)(handlers.GetUsers)(w, r)
		} else if r.Method == http.MethodDelete {
			middleware.RoleMiddleware(1, 2)(handlers.DeleteUser)(w, r)
		} else {
			http.Error(w, `{"message": "Method not allowed"}`, http.StatusMethodNotAllowed)
		}
	})))

	// CRUD PRODUCTS
	mux.HandleFunc("/api/products", middleware.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			middleware.RoleMiddleware(1)(handlers.CreateProduct)(w, r)
		} else if r.Method == http.MethodPut {
			middleware.RoleMiddleware(1, 2)(handlers.UpdateProducts)(w, r)
		} else if r.Method == http.MethodGet {
			middleware.RoleMiddleware(1, 2, 3)(handlers.GetProducts)(w, r)
		} else if r.Method == http.MethodDelete {
			middleware.RoleMiddleware(1, 2)(handlers.DeleteProduct)(w, r)
		} else {
			http.Error(w, `{"message": "Method not allowed"}`, http.StatusMethodNotAllowed)
		}
	}))

	// AI
	mux.HandleFunc("/api/generate-desc", middleware.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handlers.GenerateProductDescription(w, r)
		} else {
			http.Error(w, `{"message": "Method not allowed"}`, http.StatusMethodNotAllowed)
		}
	}))

	port := ":8080"
	log.Printf("Server berjalan di http://localhost%s\n", port)

	handlerCORS := middleware.EnableCORS(mux)

	err := http.ListenAndServe(port, handlerCORS);

	if err != nil {
		log.Fatal("Gagal menjalankan server: ", err)
	}

}