package main

import (
	"log"
	"net/http"

	"github.com/jewelmia/GoDomain/internal/application"
	appInvoice "github.com/jewelmia/GoDomain/internal/application/invoice"
	appUser "github.com/jewelmia/GoDomain/internal/application/user"
	"github.com/jewelmia/GoDomain/internal/infrastructure/db"
	repo "github.com/jewelmia/GoDomain/internal/infrastructure/persistence"
	httpRoutes "github.com/jewelmia/GoDomain/internal/interfaces/http"

	_ "github.com/jewelmia/GoDomain/internal/interfaces/http/swagger"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	// 1️⃣ Database connection
	dsn := "postgres://godomain:password@localhost:5432/go_domain?sslmode=disable"
	dbConn := db.NewPostgres(dsn)

	// 2️⃣ Repositories
	userRepo := repo.NewUserRepoInMemory(dbConn)
	invoiceRepo := repo.NewInvoiceRepoPostgres(dbConn)
	// paymentRepo := repo.NewPaymentRepoPostgres(dbConn)

	// 3️⃣ Services / UseCases
	userService := appUser.NewUserService(userRepo)
	createInvoiceUC := appInvoice.NewCreateInvoiceUseCase(invoiceRepo)
	getAllInvoiceUc := appInvoice.NewGetAllInvoiceUseCase(invoiceRepo)
	// payInvoiceUC := appInvoice.NewPayInvoiceUseCase(invoiceRepo)

	// 4️⃣ Application container
	container := &application.Container{}
	container.User.Service = userService
	container.Invoice.Create = createInvoiceUC
	container.Invoice.AllInvoice = getAllInvoiceUc
	// container.Invoice.Pay = payInvoiceUC

	// 5️⃣ HTTP Router
	mux := http.NewServeMux()
	httpRoutes.RegisterRoutes(mux, container)

	// 6️⃣ Swagger UI
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	// 7️⃣ Start server
	addr := ":7000"
	log.Printf("Server running on %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
