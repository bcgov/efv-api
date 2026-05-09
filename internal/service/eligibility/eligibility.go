package eligibility

import (
	"fmt"
	"net/http"
	"time"

	"github.com/bcgov/efv-api/internal/mocks"
	"github.com/bcgov/efv-api/internal/models"
)

// ListFactors returns available factor definitions (mocked).
func ListFactors() ([]models.FactorDefinition, int) {
	return mocks.ListFactorDefinitions(), http.StatusOK
}

// Verify validates the request and evaluates provided factors (mocked).
// Returns either an EligibilityResponse on success or a models.ErrorResponse with appropriate status code.
func Verify(req models.EligibilityRequest) (interface{}, int) {
	// validation: require at least one factor
	if len(req.Factors) == 0 {
		return models.ErrorResponse{Code: "invalid_request", Message: "at least one factor is required"}, http.StatusBadRequest
	}

	// validate each factor
	for i, f := range req.Factors {
		if f.Type == "" {
			return models.ErrorResponse{Code: "invalid_factor", Message: fmt.Sprintf("factor type is required (index: %d)", i)}, http.StatusBadRequest
		}
		if f.Value == "" {
			return models.ErrorResponse{Code: "invalid_factor", Message: "factor value is required for type: " + string(f.Type)}, http.StatusBadRequest
		}
	}

	// simple mock evaluation: mark each factor eligible and echo
	var results []models.FactorResult
	for _, f := range req.Factors {
		results = append(results, models.FactorResult{
			Eligible: true,
			Reason:   "Mock: passes",
			Type:     f.Type,
			Value:    f.Value,
		})
	}
	resp := models.EligibilityResponse{
		ApplicantID: req.ApplicantID,
		Eligible:    true,
		Factors:     results,
		Timestamp:   time.Now().UTC().Format(time.RFC3339),
	}
	return resp, http.StatusOK
}
