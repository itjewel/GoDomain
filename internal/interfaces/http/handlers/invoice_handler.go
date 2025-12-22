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
// @Param invoice body object true "Invoice Payload"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
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

func GetOneInvoiceHandler(
	uc *appInvoice.GetInvoiceByIdUseCase,
) http.HandlerFunc {	
 
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
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
