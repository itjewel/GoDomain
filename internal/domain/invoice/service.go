package invoice

import "errors"

type InvoiceService struct {
	repo InvoiceRepository
}

func NewInvoiceService(repo InvoiceRepository) *InvoiceService {
	return &InvoiceService{repo: repo}
}

// Use case: Create new invoice with some validation
func (s *InvoiceService) CreateInvoice(id, userID string, amount float64) (*Invoice, error) {
	if amount <= 0 {
		return nil, errors.New("amount must be greater than 0")
	}
	inv := NewInvoice(id, userID, amount, "pending")
	err := s.repo.Save(inv)
	if err != nil {
		return nil, err
	}
	return inv, nil
}

// Use case: Mark invoice as paid
func (s *InvoiceService) PayInvoice(id string) error {
	inv, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	inv.Status = "paid"
	return s.repo.Save(inv)
}
