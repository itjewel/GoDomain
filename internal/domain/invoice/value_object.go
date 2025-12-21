package invoice

import "errors"

// Example of a value object (immutable)
type Money struct {
	Amount   float64
	Currency string
}

func NewMoney(amount float64, currency string) (Money, error) {
	if amount < 0 {
		return Money{}, errors.New("amount cannot be negative")
	}
	if currency == "" {
		return Money{}, errors.New("currency cannot be empty")
	}
	return Money{
		Amount:   amount,
		Currency: currency,
	}, nil
}
