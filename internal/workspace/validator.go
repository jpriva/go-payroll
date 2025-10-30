package workspace

import (
	"fmt"

	"github.com/google/uuid"
)

const (
	maxCodeLength = 20
	maxNameLength = 100
)

type Validator struct {
	errs map[string]string
}

func NewValidator() *Validator {
	return &Validator{errs: make(map[string]string)}
}

func (v *Validator) Errors() map[string]string {
	return v.errs
}

func (v *Validator) HasErrors() bool {
	return len(v.errs) > 0
}

func (v *Validator) ValidateCode(code string) {
	if code == "" {
		v.errs["Code"] = "is empty"
	} else if len(code) > maxCodeLength {
		v.errs["Code"] = fmt.Sprintf("must be less than %d characters", maxCodeLength)
	}
}

func (v *Validator) ValidateName(name string) {
	if name == "" {
		v.errs["Name"] = "is empty"
	} else if len(name) > maxNameLength {
		v.errs["Name"] = fmt.Sprintf("must be less than %d characters", maxNameLength)
	}
}

func (v *Validator) ValidateStatus(status *WorkspaceStatus) {
	if status != nil && !status.IsValid() {
		v.errs["Status"] = "is invalid"
	}
}

func (v *Validator) ValidateCountryID(countryID uuid.UUID) {
	if countryID == uuid.Nil {
		v.errs["CountryID"] = "is empty"
	}
}
