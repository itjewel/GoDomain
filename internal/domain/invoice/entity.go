package invoice

type Invoice struct {
	ID     string  `json:"id"`
	UserID string  `json:"user_id"`
	Amount float64 `json:"amount"`
	Status string  `json:"status"`
}

// Constructor
func NewInvoice(id, userID string, amount float64, status string) *Invoice {
	return &Invoice{
		ID:     id,
		UserID: userID,
		Amount: amount,
		Status: status,
	}
}

// Example of business logic inside entity
func (i *Invoice) IsPaid() bool {
	return i.Status == "paid"
}