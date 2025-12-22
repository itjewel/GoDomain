package application

import (
	appInvoice "github.com/jewelmia/GoDomain/internal/application/invoice"
	appUser "github.com/jewelmia/GoDomain/internal/application/user"
)

type Container struct {
	User struct {
		Service *appUser.UserService
	}

	Invoice struct {
		Create *appInvoice.CreateInvoiceUseCase
		AllInvoice *appInvoice.GetAllInvoiceUseCase
		Pay    *appInvoice.PayInvoiceUseCase
	}
}
