package invoice

import domain "github.com/jewelmia/GoDomain/internal/domain/invoice"

type GetAllInvoiceUseCase struct {
	repo domain.InvoiceRepository
}

func NewGetAllInvoiceUseCase(
	repo domain.InvoiceRepository,
) *GetAllInvoiceUseCase{
	return &GetAllInvoiceUseCase{repo:repo}
}

func (uc *GetAllInvoiceUseCase) Execute()([]*domain.Invoice, error){
return uc.repo.FindAll()
}