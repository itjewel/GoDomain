package invoice

import (
	"errors"

	domain "github.com/jewelmia/GoDomain/internal/domain/invoice"
)

type CreateInvoiceUseCase struct {
	repo domain.InvoiceRepository
}

func NewCreateInvoiceUseCase(
	repo domain.InvoiceRepository,
) *CreateInvoiceUseCase {
	return &CreateInvoiceUseCase{repo: repo}
}

type CreateInvoiceCommand struct {
	ID     string
	UserID string
	Amount float64
}

func (uc *CreateInvoiceUseCase) Execute(
	cmd CreateInvoiceCommand,
) (*domain.Invoice, error) {

	if cmd.Amount <= 0 {
		return nil, errors.New("amount must be greater than 0")
	}

inv, err := domain.NewInvoice(cmd.ID, cmd.UserID, cmd.Amount)
if err != nil {
	return nil, err
}


	return inv, nil
}
