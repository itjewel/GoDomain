package invoice

// Domain layer এ interface define করা হয়
type InvoiceRepository interface {
	Save(inv *Invoice) error
	GetByID(id string) (*Invoice, error)
}
