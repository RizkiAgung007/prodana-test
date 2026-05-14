package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
)

type ContextKey string

const UserIDKey ContextKey = "user_id"
const RoleIDKey ContextKey = "role_id"

// MEmastikan request token jwt valid
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// Mengambl header authhoziation
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, `{"message": "Unauthorized, token tidak ditemukan"}`, http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		secretKey := os.Getenv("JWT_SECRET")

		// Melakukan validasi token
		token , err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, `{"message": "Unauthorized, token tidak valid atau exp"}`, http.StatusUnauthorized)
			return
		}

		// Estrak data dari token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, `{"message": "Unauthorized, gagal membaca token"}`, http.StatusUnauthorized)
			return
		}

		userID := uint(claims["user_id"].(float64))
		roleID := uint(claims["role_id"].(float64))

		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		ctx = context.WithValue(ctx, RoleIDKey, roleID)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

// Membatasi akses berdasarkan role id
func RoleMiddleware(allowedRoles ...uint) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")

			// MEngambil role_id dari context AuthMiddleware
			roleIDValue := r.Context().Value(RoleIDKey)
			if roleIDValue == nil {
				http.Error(w, `{"message": "Forbidded, anda tidak memillki akses"}`, http.StatusForbidden)
				return
			}

			userRole := roleIDValue.(uint)
			isAllowed := false

			// Cek role user apakah diizinkan didalam daftar atau tidak
			for _, role := range allowedRoles {
				if userRole == role {
					isAllowed = true
					break
				}
			}

			if !isAllowed {
				http.Error(w, `{"message": "Forbiddedn", akses ditolak untuk role ini}`, http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		}
	}
}