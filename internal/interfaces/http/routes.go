package http

import (
	"net/http"

	appUser "github.com/jewelmia/GoDomain/internal/application/user"
	"github.com/jewelmia/GoDomain/internal/interfaces/http/handlers"
)

func RegisterRoutes(mux *http.ServeMux, userService *appUser.UserService) {
	mux.HandleFunc("/users", handlers.CreateUserHandler(userService))
	mux.HandleFunc("/users/get", handlers.GetUserHandler(userService))
}
