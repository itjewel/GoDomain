// internal/interfaces/http/handlers/dto.go
package handlers

type InvoiceResponse struct {
	ID     string  `json:"id" example:"1"`
	UserID string  `json:"user_id" example:"123"`
	Amount float64 `json:"amount" example:"1500"`
	Status string  `json:"status" example:"pending"`
}
