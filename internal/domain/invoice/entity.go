package invoice

import "errors"

/*
Domain constants (no magic strings)
*/
type Status string

const (
	StatusPending Status = "pending"
	StatusPaid    Status = "paid"
)

type Invoice struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	Amount float64 `json:"amount"`
	Status Status `json:"status"`
}

// Constructor (factory)
func NewInvoice( userID string, amount float64) (*Invoice, error) {

	if userID == "" {
		return nil, errors.New("user id is required")
	}
	if amount <= 0 {
		return nil, errors.New("amount must be greater than zero")
	}

	return &Invoice{
		// ID:     id,
		UserID: userID,
		Amount: amount,
		Status: StatusPending,
	}, nil
}

// Query method
func (i *Invoice) IsPaid() bool {
	return i.Status == StatusPaid
}

// Command method (state transition)
func (i *Invoice) MarkAsPaid() error {
	if i.Status == StatusPaid {
		return errors.New("invoice already paid")
	}
	i.Status = StatusPaid
	return nil
}
