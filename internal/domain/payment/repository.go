package payment

type PaymentRepository interface {
	Save(p *Payment) error
	
	GetByID(id string) (*Payment, error)
	GetByInvoiceID(invoiceID string) ([]*Payment, error)
}
