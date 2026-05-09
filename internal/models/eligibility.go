package models

type Factor struct {
	Type  FactorType `json:"type"`
	Value string     `json:"value"`
}

type FactorDefinition struct {
	Description string     `json:"description,omitempty"`
	Name        string     `json:"name,omitempty"`
	Required    bool       `json:"required,omitempty"`
	Type        FactorType `json:"type,omitempty"`
}

// FactorType represents the allowed factor type enum (string-backed).
type FactorType string

type FactorResult struct {
	Eligible bool       `json:"eligible"`
	Reason   string     `json:"reason,omitempty"`
	Type     FactorType `json:"type"`
	Value    string     `json:"value"`
}

// Eligibility request/response
type EligibilityRequest struct {
	ApplicantID string   `json:"applicantId,omitempty"`
	Factors     []Factor `json:"factors,omitempty"`
}

type EligibilityResponse struct {
	ApplicantID string         `json:"applicantId,omitempty"`
	Eligible    bool           `json:"eligible"`
	Factors     []FactorResult `json:"factors,omitempty"`
	Timestamp   string         `json:"timestamp,omitempty"`
}

type HealthResponse struct {
	Status string `json:"status"`
}
