package handlers

import (
	"encoding/json"
	"net/http"

	appInvoice "github.com/jewelmia/GoDomain/internal/application/invoice"
	response "github.com/jewelmia/GoDomain/internal/interfaces/http/response"
)

// JSONResponse helper

// CreateInvoiceHandler godoc
// @Summary Create Invoice
// @Description Create a new invoice
// @Tags Invoice
// @Accept json
// @Produce json
// @Param invoice body CreateInvoiceRequest true "Invoice Payload"
// @Success 201 {object} invoice.Invoice
// @Failure 400 {object} response.InvoiceResponse
// @Router /invoice/create [post]

func CreateInvoiceHandler(
	uc *appInvoice.CreateInvoiceUseCase,
) http.HandlerFunc {
	
 
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			response.JSONResponse(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
			return
		}
		defer r.Body.Close()

		var req struct {
			UserID string  `json:"user_id"`
			Amount float64 `json:"amount"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			response.JSONResponse(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
			return
		}

		inv, err := uc.Execute(appInvoice.CreateInvoiceCommand{
			UserID: req.UserID,
			Amount: req.Amount,
		})

		if err != nil {
			response.JSONResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		response.JSONResponse(w, http.StatusCreated, inv)
	}
		
}


// GetAllInvoiceHandler godoc
// @Summary All Invoice
// @Description show all invoice
// @Tags Invoice
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Router /invoices [get]
func GetAllInvoiceHandler(uc *appInvoice.GetAllInvoiceUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		invoices, err := uc.Execute()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(invoices)
	}
}


// GetOneInvoiceHandler godoc
// @Summary Get invoice by ID
// @Description Retrieve a single invoice by its ID
// @Tags Invoice
// @Produce json
// @Param id path string true "Invoice ID"
// @Success 200 {object} invoice.Invoice
// @Failure 400 {object} response.InvoiceResponse
// @Router /invoices/{id} [get]
func GetOneInvoiceHandler(
	uc *appInvoice.GetInvoiceByIdUseCase,
) http.HandlerFunc {	
 
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			response.JSONResponse(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
			return
		}
		id := r.PathValue("id")
		
		if id == "" {
			response.JSONResponse(w, http.StatusBadRequest, map[string]string{"error": "invoice id required"})
			return
		}
	   invoice, err := uc.Execute(id)

		if err != nil {
			response.JSONResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		response.JSONResponse(w, http.StatusCreated, invoice)
	}
		
}
