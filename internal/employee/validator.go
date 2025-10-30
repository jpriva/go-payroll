package employee

import (
	"fmt"
	"net/mail"
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

func (v *Validator) ValidateFirstName(firstName string) {
	if firstName == "" {
		v.errs["FirstName"] = "is empty"
	} else if len(firstName) > maxFirstNameLength {
		v.errs["FirstName"] = fmt.Sprintf("must be less than %d characters", maxFirstNameLength)
	}
}

func (v *Validator) ValidateLastName(lastName string) {
	if lastName == "" {
		v.errs["LastName"] = "is empty"
	} else if len(lastName) > maxLastNameLength {
		v.errs["LastName"] = fmt.Sprintf("must be less than %d characters", maxLastNameLength)
	}
}

func (v *Validator) ValidateEmail(email string) {
	if email == "" {
		v.errs["Email"] = "is empty"
		return
	}
	if len(email) > maxEmailLength {
		v.errs["Email"] = fmt.Sprintf("must be less than %d characters", maxEmailLength)
	}
	if _, err := mail.ParseAddress(email); err != nil {
		v.errs["Email"] = "is not a valid email format"
	}
}

func (v *Validator) ValidateDocTypeID(docTypeID uuid.UUID) {
	if docTypeID == uuid.Nil {
		v.errs["DocTypeID"] = "is empty"
	}
}

func (v *Validator) ValidateDocNumber(docNumber string) {
	if docNumber == "" {
		v.errs["DocNumber"] = "is empty"
	} else if len(docNumber) > maxDocNumberLength {
		v.errs["DocNumber"] = fmt.Sprintf("must be less than %d characters", maxDocNumberLength)
	}
}

func (v *Validator) ValidateBirthDate(birthDate *time.Time) {
	if birthDate != nil && birthDate.After(time.Now()) {
		v.errs["BirthDate"] = "cannot be in the future"
	}
}

func (v *Validator) ValidateGender(gender *string) {
	if gender != nil {
		g := EmployeeGender(*gender)
		if !g.IsValid() {
			v.errs["Gender"] = "is invalid"
		}
	}
}

func (v *Validator) ValidatePhone(phone *string) {
	if phone != nil && len(*phone) > maxPhoneLength {
		v.errs["Phone"] = fmt.Sprintf("must be less than %d characters", maxPhoneLength)
	}
}

func (v *Validator) ValidateAddress(address *string) {
	if address != nil && len(*address) > maxAddressLength {
		v.errs["Address"] = fmt.Sprintf("must be less than %d characters", maxAddressLength)
	}
}
