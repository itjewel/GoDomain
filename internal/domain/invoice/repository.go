package invoice

type InvoiceRepository interface {
	Save(inv *Invoice) error
	GetByID(id string) (*Invoice, error)
}
