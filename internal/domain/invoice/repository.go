package invoice

type InvoiceRepository interface {
	Save(inv *Invoice) error
	FindAll()([]*Invoice, error)
	GetByID(id string) (*Invoice, error)
}
