package persistence

import (
	"database/sql"
	"errors"

	"github.com/jewelmia/GoDomain/internal/domain/invoice"
)

type PaymentRepoPostgres struct {
	db *sql.DB
}

func NewPaymentRepoPostgres(db *sql.DB) invoice.InvoiceRepository {
	return &InvoiceRepoPostgres{db: db}
}

func (r *InvoiceRepoPostgres) Save(inv *invoice.Invoice) error {
	query := `
	INSERT INTO invoices (id, user_id, amount, status)
	VALUES ($1, $2, $3, $4)
	ON CONFLICT (id) DO UPDATE
	SET user_id=EXCLUDED.user_id, amount=EXCLUDED.amount, status=EXCLUDED.status;
	`
	_, err := r.db.Exec(query, inv.ID, inv.UserID, inv.Amount, inv.Status)
	return err
}

func (r *InvoiceRepoPostgres) GetByID(id string) (*invoice.Invoice, error) {
	inv := &invoice.Invoice{}
	row := r.db.QueryRow(`SELECT id, user_id, amount, status FROM invoices WHERE id=$1`, id)
	err := row.Scan(&inv.ID, &inv.UserID, &inv.Amount, &inv.Status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("invoice not found")
		}
		return nil, err
	}
	return inv, nil
}
