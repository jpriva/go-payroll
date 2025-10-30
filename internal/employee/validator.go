package employee

import (
	"fmt"
	"net/mail"
	"payroll/internal/platform/validation"
	"time"

	"github.com/google/uuid"
)

const (
	maxFirstNameLength = 50
	maxLastNameLength  = 50
	maxEmailLength     = 100
	maxAddressLength   = 200
	maxDocNumberLength = 50
	maxPhoneLength     = 20
)

type Validator struct {
	validation.Validator
}

func NewValidator() *Validator {
	return &Validator{*validation.New()}
}

func (v *Validator) ValidateFirstName(firstName string) {
	if firstName == "" {
		v.AddError("FirstName", "is empty")
	} else if len(firstName) > maxFirstNameLength {
		v.AddError("FirstName", fmt.Sprintf("must be less than %d characters", maxFirstNameLength))
	}
}

func (v *Validator) ValidateLastName(lastName string) {
	if lastName == "" {
		v.AddError("LastName", "is empty")
	} else if len(lastName) > maxLastNameLength {
		v.AddError("LastName", fmt.Sprintf("must be less than %d characters", maxLastNameLength))
	}
}

func (v *Validator) ValidateEmail(email string) {
	if email == "" {
		v.AddError("Email", "is empty")
		return
	}
	if len(email) > maxEmailLength {
		v.AddError("Email", fmt.Sprintf("must be less than %d characters", maxEmailLength))
	}
	if _, err := mail.ParseAddress(email); err != nil {
		v.AddError("Email", "is not a valid email format")
	}
}

func (v *Validator) ValidateDocTypeID(docTypeID uuid.UUID) {
	if docTypeID == uuid.Nil {
		v.AddError("DocTypeID", "is empty")
	}
}

func (v *Validator) ValidateDocNumber(docNumber string) {
	if docNumber == "" {
		v.AddError("DocNumber", "is empty")
	} else if len(docNumber) > maxDocNumberLength {
		v.AddError("DocNumber", fmt.Sprintf("must be less than %d characters", maxDocNumberLength))
	}
}

func (v *Validator) ValidateBirthDate(birthDate *time.Time) {
	if birthDate != nil && birthDate.After(time.Now()) {
		v.AddError("BirthDate", "cannot be in the future")
	}
}

func (v *Validator) ValidateGender(gender *string) {
	if gender != nil {
		g := EmployeeGender(*gender)
		if !g.IsValid() {
			v.AddError("Gender", "is invalid")
		}
	}
}

func (v *Validator) ValidatePhone(phone *string) {
	if phone != nil && len(*phone) > maxPhoneLength {
		v.AddError("Phone", fmt.Sprintf("must be less than %d characters", maxPhoneLength))
	}
}

func (v *Validator) ValidateAddress(address *string) {
	if address != nil && len(*address) > maxAddressLength {
		v.AddError("Address", fmt.Sprintf("must be less than %d characters", maxAddressLength))
	}
}
