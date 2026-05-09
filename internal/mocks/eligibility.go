package mocks

import (
	"github.com/bcgov/efv-api/internal/models"
)

// ListFactorDefinitions returns a small set of mock factor definitions.
func ListFactorDefinitions() []models.FactorDefinition {
	return []models.FactorDefinition{
		{
			Description: "Verifies if the applicant meets minimum age",
			Name:        "Age Verification",
			Required:    true,
			Type:        "age",
		},
		{
			Description: "Verifies residency status",
			Name:        "Residency Verification",
			Required:    false,
			Type:        "residence",
		},
	}
}
