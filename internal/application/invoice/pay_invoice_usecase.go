package invoice

import (
	domain "github.com/jewelmia/GoDomain/internal/domain/invoice"
)

type PayInvoiceUseCase struct {
	repo domain.InvoiceRepository
}

func NewPayInvoiceUseCase(
	repo domain.InvoiceRepository,
) *PayInvoiceUseCase {
	return &PayInvoiceUseCase{repo: repo}
}

func (uc *PayInvoiceUseCase) Execute(id string) error {
	inv, err := uc.repo.GetByID(id)
	if err != nil {
		return err
	}

	inv.MarkAsPaid() // better than inv.Status = "paid"
	return uc.repo.Save(inv)
}
