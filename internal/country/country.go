package country

import (
	"context"
	"payroll/internal/apperror"
	"payroll/internal/domain"
	"strings"

	"github.com/google/uuid"
)

type Country struct {
	domain.BaseEntity
	ID         uuid.UUID
	Code       string
	Name       string
	CoinCode   string
	CoinSymbol string
}

type CreateCountryParams struct {
	Code       string
	Name       string
	CoinCode   string
	CoinSymbol string
}

type UpdateCountryParams struct {
	Code       *string
	Name       *string
	CoinCode   *string
	CoinSymbol *string
}

func NewCountry(params CreateCountryParams) (*Country, error) {
	validator := NewValidator()

	params.Code = strings.TrimSpace(params.Code)
	params.Name = strings.TrimSpace(params.Name)
	params.CoinCode = strings.TrimSpace(params.CoinCode)
	params.CoinSymbol = strings.TrimSpace(params.CoinSymbol)

	validator.ValidateCode(params.Code)
	validator.ValidateName(params.Name)
	validator.ValidateCoinCode(params.CoinCode)
	validator.ValidateCoinSymbol(params.CoinSymbol)

	if validator.HasErrors() {
		return nil, apperror.NewValidationError("Country", validator.Errors())
	}

	country := &Country{
		Code:       params.Code,
		Name:       params.Name,
		CoinCode:   params.CoinCode,
		CoinSymbol: params.CoinSymbol,
	}
	country.Initialize()

	return country, nil
}

type Repository interface {
	Create(ctx context.Context, country *Country) error
	Update(ctx context.Context, country *Country) error
	Delete(ctx context.Context, id uuid.UUID) error
	ExistsByCode(ctx context.Context, code string) (bool, error)
	GetByID(ctx context.Context, id uuid.UUID) (*Country, error)
	GetByCode(ctx context.Context, code string) (*Country, error)
	ListAll(ctx context.Context) ([]*Country, error)
}
