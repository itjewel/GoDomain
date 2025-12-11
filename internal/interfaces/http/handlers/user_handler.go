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

// CreateUserHandler godoc
// @Summary Create a new user
// @Description Creates a new user with ID, name, email
// @Tags User
// @Accept json
// @Produce json
// @Param user body object true "User Payload"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /users/create [post]
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

// GetUserHandler godoc
// @Summary Get user by ID
// @Description Fetch a user by query parameter id
// @Tags User
// @Produce json
// @Param id query string true "User ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /users/get [get]
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
