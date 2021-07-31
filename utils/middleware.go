package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pbutarbutar/crud-golang/database"
	"github.com/pbutarbutar/crud-golang/entity"
)

type UserKey struct{}

//MiddlewareValidateRefreshToken is
func MiddlewareValidateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		token, _ := extractToken(r)
		claims, is_ok := extractClaims(token)

		fmt.Println(claims["user_id"])
		apiResp := entity.ApiResponse{}
		if !is_ok {
			w.WriteHeader(http.StatusUnauthorized)
			apiResp.Status = http.StatusUnauthorized
			apiResp.Success = false
			apiResp.Message = "Authorization Failed"
			json.NewEncoder(w).Encode(apiResp)
			return
		}

		var users entity.Users
		if err := database.Connector.First(&users, "user_id = ? ", claims["user_id"]).Error; err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			apiResp.Status = http.StatusUnauthorized
			apiResp.Success = false
			apiResp.Message = "Login Failed - Your Authorization not found!"
			json.NewEncoder(w).Encode(apiResp)
			return
		}

		ctx := context.WithValue(r.Context(), UserKey{}, users)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
