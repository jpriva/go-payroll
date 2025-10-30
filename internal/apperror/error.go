package apperror

import (
	"encoding/json"
	"fmt"
)

type Type string

const (
	TypeNotFound  Type = "NOT_FOUND"
	TypeInvalid   Type = "INVALID_INPUT"
	TypeDuplicate Type = "DUPLICATE_ENTRY"
)

type DomainError struct {
	Type    Type              `json:"type"`
	Origin  string            `json:"origin"`
	Message string            `json:"message"`
	Details map[string]string `json:"details,omitempty"`
}

func (e *DomainError) Error() string {
	if len(e.Details) > 0 {
		detailBytes, err := json.Marshal(e.Details)
		if err != nil {
			return fmt.Sprintf("[%s/%s]: %s. Details: %s", e.Origin, e.Type, e.Message, string(detailBytes))
		}
	}

	return fmt.Sprintf("[%s/%s]: %s", e.Origin, e.Type, e.Message)
}

func New(errType Type, origin, msg string) error {
	return &DomainError{
		Type:    errType,
		Origin:  origin,
		Message: msg,
	}
}

func NewValidationError(origin string, details map[string]string) error {
	return &DomainError{
		Type:    TypeInvalid,
		Origin:  origin,
		Message: "Validation failed",
		Details: details,
	}
}
