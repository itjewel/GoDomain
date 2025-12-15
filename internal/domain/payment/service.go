package payment

import "errors"

type PaymentService struct {
	repo PaymentRepository
}

func NewPaymentService(repo PaymentRepository) *PaymentService {
	return &PaymentService{repo: repo}
}

// Create new payment
func (s *PaymentService) CreatePayment(id, invoiceID, userID string, amount float64, method string) (*Payment, error) {
	if amount <= 0 {
		return nil, errors.New("amount must be greater than 0")
	}
	p := NewPayment(id, invoiceID, userID, amount, method)
	err := s.repo.Save(p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// Complete payment
func (s *PaymentService) CompletePayment(id string) error {
	p, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	p.Status = "completed"
	return s.repo.Save(p)
}

// Get payments for an invoice
func (s *PaymentService) PaymentsByInvoice(invoiceID string) ([]*Payment, error) {
	return s.repo.GetByInvoiceID(invoiceID)
}
