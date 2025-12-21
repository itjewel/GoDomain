package invoice

import "errors"

type Invoice struct {
	ID     string  `json:"id"`
	UserID string  `json:"user_id"`
	Amount float64 `json:"amount"`
	Status string  `json:"status"`
}

// Constructor
func NewInvoice(id, userID string, amount float64) *Invoice {
	return &Invoice{
		ID:     id,
		UserID: userID,
		Amount: amount,
		Status: "pending", // default status
	}
}

// Business rule: check if invoice is paid
func (i *Invoice) IsPaid() bool {
	return i.Status == "paid"
}

// Business rule: mark invoice as paid
func (i *Invoice) MarkPaid() error {
	if i.Status == "paid" {
		return errors.New("invoice already paid")
	}
	i.Status = "paid"
	return nil
}
