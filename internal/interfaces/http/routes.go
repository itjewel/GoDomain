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
	mux.HandleFunc("/users", handlers.CreateUserHandler(c.User.Service))
	mux.HandleFunc("/users/get", handlers.GetUserHandler(c.User.Service))

	// Invoice routes
	mux.HandleFunc(
		"/invoice/create",
		handlers.CreateInvoiceHandler(c.Invoice.Create),
	)
}
