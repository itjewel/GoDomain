package payment

type Payment struct {
	ID        string  `json:"id"`
	InvoiceID string  `json:"invoice_id"`
	UserID    string  `json:"user_id"`
	Amount    float64 `json:"amount"`
	Status    string  `json:"status"` // pending, completed, failed
	Method    string  `json:"method"` // e.g., credit_card, bKash
}

// Constructor
func NewPayment(id, invoiceID, userID string, amount float64, method string) *Payment {
	return &Payment{
		ID:        id,
		InvoiceID: invoiceID,
		UserID:    userID,
		Amount:    amount,
		Status:    "pending",
		Method:    method,
	}
}

// Business logic example
func (p *Payment) IsCompleted() bool {
	return p.Status == "completed"
}
