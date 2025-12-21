package payment

import (
	"errors"

	domain "github.com/jewelmia/GoDomain/internal/domain/payment"
)


type PaymentService struct {
	repo domain.PaymentRepository
}

func NewPaymentService(repo domain.PaymentRepository) *PaymentService {
	return &PaymentService{repo: repo}
}

// Create new payment
func (s *PaymentService) CreatePayment(id, invoiceID, userID string, amount float64, method string) (*domain.Payment, error) {
	if amount <= 0 {
		return nil, errors.New("amount must be greater than 0")
	}
	p := domain.NewPayment(id, invoiceID, userID, amount, method)
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
func (s *PaymentService) PaymentsByInvoice(invoiceID string) ([]*domain.Payment, error) {
	return s.repo.GetByInvoiceID(invoiceID)
}
