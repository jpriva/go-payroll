package country

import (
	"fmt"
	"payroll/internal/platform/validation"
)

const (
	maxCodeLength       = 3
	maxNameLength       = 100
	maxCoinCodeLength   = 3
	maxCoinSymbolLength = 5
)

type Validator struct {
	validation.Validator
}

func NewValidator() *Validator {
	return &Validator{*validation.New()}
}

func (v *Validator) ValidateCode(code string) {
	if code == "" {
		v.AddError("Code", "is empty")
	} else if len(code) > maxCodeLength {
		v.AddError("Code", fmt.Sprintf("must be less than %d characters", maxCodeLength))
	}
}

func (v *Validator) ValidateName(name string) {
	if name == "" {
		v.AddError("Name", "is empty")
	} else if len(name) > maxNameLength {
		v.AddError("Name", fmt.Sprintf("must be less than %d characters", maxNameLength))
	}
}

func (v *Validator) ValidateCoinCode(coinCode string) {
	if coinCode == "" {
		v.AddError("CoinCode", "is empty")
	} else if len(coinCode) > maxCoinCodeLength {
		v.AddError("CoinCode", fmt.Sprintf("must be less than %d characters", maxCoinCodeLength))
	}
}
func (v *Validator) ValidateCoinSymbol(coinSymbol string) {
	if coinSymbol == "" {
		v.AddError("CoinSymbol", "is empty")
	} else if len(coinSymbol) > maxCoinSymbolLength {
		v.AddError("CoinSymbol", fmt.Sprintf("must be less than %d characters", maxCoinSymbolLength))
	}
}
