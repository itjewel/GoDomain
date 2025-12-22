package http

import (
	"net/http"

	"github.com/jewelmia/GoDomain/internal/application"
	handlers "github.com/jewelmia/GoDomain/internal/interfaces/http/handlers"
)

func RegisterRoutes(
	mux *http.ServeMux,
	c *application.Container,
) {
	// User routes
	mux.HandleFunc("POST /users", handlers.CreateUserHandler(c.User.Service)) // POST
	mux.HandleFunc("GET /users/get", handlers.GetUserHandler(c.User.Service)) // GET

	// Invoice routes
	mux.HandleFunc(
		"POST /invoice/create",
		handlers.CreateInvoiceHandler(c.Invoice.Create)) // POST
	// GET /invoices
	mux.HandleFunc("/invoices", handlers.GetAllInvoiceHandler(c.Invoice.AllInvoice))
	mux.HandleFunc(
		"GET /invoice/{id}",
		handlers.CreateInvoiceHandler(c.Invoice.Create)) // GET
}
