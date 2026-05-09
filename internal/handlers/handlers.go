package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/bcgov/efv-api/internal/models"
	eligibilitySvc "github.com/bcgov/efv-api/internal/service/eligibility"
)

func writeJSON(w http.ResponseWriter, code int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}

// RegisterRoutes registers HTTP handlers on the provided mux.
func RegisterRoutes(mux *http.ServeMux) {

	// Eligibility endpoints (from OpenAPI spec)
	mux.HandleFunc("/api/v1/eligibility/factors", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			f, _ := eligibilitySvc.ListFactors()
			writeJSON(w, http.StatusOK, f)
		default:
			writeJSON(w, http.StatusMethodNotAllowed, models.ErrorResponse{Code: "method_not_allowed", Message: "method not allowed"})
		}
	})

	mux.HandleFunc("/api/v1/eligibility/verify", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			var req models.EligibilityRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				writeJSON(w, http.StatusBadRequest, models.ErrorResponse{Code: "invalid_payload", Message: "invalid JSON"})
				return
			}
			resp, status := eligibilitySvc.Verify(req)
			writeJSON(w, status, resp)
		default:
			writeJSON(w, http.StatusMethodNotAllowed, models.ErrorResponse{Code: "method_not_allowed", Message: "method not allowed"})
		}
	})

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			writeJSON(w, http.StatusMethodNotAllowed, models.ErrorResponse{Code: "method_not_allowed", Message: "method not allowed"})
			return
		}
		writeJSON(w, http.StatusOK, models.HealthResponse{Status: "ok"})
	})
}
