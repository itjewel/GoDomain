package invoice

import (
	"errors"

	domain "github.com/jewelmia/GoDomain/internal/domain/invoice"
)

type GetInvoiceByIdUseCase struct{
	repo domain.InvoiceRepository
}

func NewGetINvoiceByIdUseCase( repo domain.InvoiceRepository) *GetInvoiceByIdUseCase{
	return &GetInvoiceByIdUseCase{repo: repo}

}

func (uc *GetInvoiceByIdUseCase) Execute(id string)(*domain.Invoice, error){

	if(id == ""){
		return nil, errors.New("Invoice id is required")
	}
	return uc.repo.GetByID(id)
}