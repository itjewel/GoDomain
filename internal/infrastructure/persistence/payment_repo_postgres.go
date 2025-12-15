package persistence

import (
	"database/sql"
	"errors"

	"github.com/jewelmia/GoDomain/internal/domain/payment"
)

// Postgres implementation of PaymentRepository
type PaymentRepoPostgres struct {
	db *sql.DB
}

// Constructor
func NewPaymentRepoPostgres(db *sql.DB) payment.PaymentRepository {
	return &PaymentRepoPostgres{db: db}
}

// Save payment
func (r *PaymentRepoPostgres) Save(p *payment.Payment) error {
	query := `
	INSERT INTO payments (id, invoice_id, user_id, amount, status, method)
	VALUES ($1, $2, $3, $4, $5, $6)
	ON CONFLICT (id) DO UPDATE
	SET invoice_id = EXCLUDED.invoice_id,
	    user_id = EXCLUDED.user_id,
	    amount = EXCLUDED.amount,
	    status = EXCLUDED.status,
	    method = EXCLUDED.method;
	`
	_, err := r.db.Exec(query, p.ID, p.InvoiceID, p.UserID, p.Amount, p.Status, p.Method)
	return err
}

// Get payment by ID
func (r *PaymentRepoPostgres) GetByID(id string) (*payment.Payment, error) {
	p := &payment.Payment{}
	row := r.db.QueryRow(`SELECT id, invoice_id, user_id, amount, status, method FROM payments WHERE id=$1`, id)
	err := row.Scan(&p.ID, &p.InvoiceID, &p.UserID, &p.Amount, &p.Status, &p.Method)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("payment not found")
		}
		return nil, err
	}
	return p, nil
}

// Get payments by InvoiceID
func (r *PaymentRepoPostgres) GetByInvoiceID(invoiceID string) ([]*payment.Payment, error) {
	rows, err := r.db.Query(`SELECT id, invoice_id, user_id, amount, status, method FROM payments WHERE invoice_id=$1`, invoiceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []*payment.Payment
	for rows.Next() {
		p := &payment.Payment{}
		if err := rows.Scan(&p.ID, &p.InvoiceID, &p.UserID, &p.Amount, &p.Status, &p.Method); err != nil {
			return nil, err
		}
		payments = append(payments, p)
	}
	return payments, nil
}
