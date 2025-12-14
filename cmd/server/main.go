package main

import (
	"log"
	"net/http"

	appUser "github.com/jewelmia/GoDomain/internal/application/user"
	"github.com/jewelmia/GoDomain/internal/infrastructure/db"
	repo "github.com/jewelmia/GoDomain/internal/infrastructure/persistence"
	httpRoutes "github.com/jewelmia/GoDomain/internal/interfaces/http"

	_ "github.com/jewelmia/GoDomain/internal/interfaces/http/swagger"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	dns := "postgres://admin:admin123@localhost:5432/app?sslmode=disable"
	dbConn := db.NewPostgres(dns)
	// DB connection
	// postgresDB, err := db.NewPostgres(dns)
	
	// Repository
	userRepo := repo.NewUserRepoInMemory(dbConn)
	// invoiceRepo := repo.InvoiceRepoPostgres(dbConn)
	// paymentRepo := repo.PaymentRepoPostgres(dbConn)

	// service
	userService := appUser.NewUserService(userRepo)
	// userInvoice := appUser.(userRepo)
	// userPayment := appUser.NewUserService(userRepo)

	mux := http.NewServeMux()
	httpRoutes.RegisterRoutes(mux, userService)
	// Serve Swagger UI at /swagger/
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	addr := ":7000"
	log.Printf("Server running on %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
