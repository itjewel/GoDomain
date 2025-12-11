package handlers

import (
	"encoding/json"
	"net/http"

	appUser "github.com/jewelmia/GoDomain/internal/application/user"
)

func JSONResponse(w http.ResponseWriter, code int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}

func CreateUserHandler(service *appUser.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			ID    string `json:"id"`
			Name  string `json:"name"`
			Email string `json:"email"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			JSONResponse(w, http.StatusBadRequest, map[string]string{"error": "invalid body"})
			return
		}
		u, err := service.CreateUser(req.ID, req.Name, req.Email)
		if err != nil {
			JSONResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		JSONResponse(w, http.StatusOK, u)
	}
}

func GetUserHandler(service *appUser.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			JSONResponse(w, http.StatusBadRequest, map[string]string{"error": "id required"})
			return
		}
		u, err := service.GetUser(id)
		if err != nil {
			JSONResponse(w, http.StatusNotFound, map[string]string{"error": err.Error()})
			return
		}
		JSONResponse(w, http.StatusOK, u)
	}
}
