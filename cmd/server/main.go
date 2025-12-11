package main

import (
	"log"
	"net/http"

	appUser "github.com/jewelmia/GoDomain/internal/application/user"
	repo "github.com/jewelmia/GoDomain/internal/infrastructure/persistence"
	httpRoutes "github.com/jewelmia/GoDomain/internal/interfaces/http"
)

func main() {
	userRepo := repo.NewUserRepoInMemory()
	userService := appUser.NewUserService(userRepo)

	mux := http.NewServeMux()
	httpRoutes.RegisterRoutes(mux, userService)

	addr := ":7000"
	log.Printf("Server running on %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
