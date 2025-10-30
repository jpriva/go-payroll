package country

import (
	"context"
	"strings"

	"payroll/internal/apperror"

	"github.com/google/uuid"
)

type Service struct {
	countryRepository Repository
}

func NewService(cr Repository) *Service {
	return &Service{
		countryRepository: cr,
	}
}

func (s *Service) ListAllCountries(ctx context.Context) ([]*Country, error) {
	return s.countryRepository.ListAll(ctx)
}

func (s *Service) GetCountryByID(ctx context.Context, id uuid.UUID) (*Country, error) {
	return s.countryRepository.GetByID(ctx, id)
}

func (s *Service) GetCountryByCode(ctx context.Context, code string) (*Country, error) {
	return s.countryRepository.GetByCode(ctx, code)
}

func (s *Service) CreateCountry(ctx context.Context, params CreateCountryParams) (*Country, error) {
	country, err := NewCountry(params)
	if err != nil {
		return nil, err
	}

	exists, err := s.countryRepository.ExistsByCode(ctx, country.Code)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, apperror.New(apperror.TypeDuplicate, "CountryService", "A country with this code already exists")
	}

	if err := s.countryRepository.Create(ctx, country); err != nil {
		return nil, err
	}

	return country, nil
}

func (s *Service) UpdateCountry(ctx context.Context, id uuid.UUID, params UpdateCountryParams) (*Country, error) {
	country, err := s.countryRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	validator := NewValidator()

	if params.Code != nil {
		trimmedCode := strings.TrimSpace(*params.Code)
		validator.ValidateCode(trimmedCode)
		country.Code = trimmedCode
	}
	if params.Name != nil {
		trimmedName := strings.TrimSpace(*params.Name)
		validator.ValidateName(trimmedName)
		country.Name = trimmedName
	}
	if params.CoinCode != nil {
		trimmedCoinCode := strings.TrimSpace(*params.CoinCode)
		validator.ValidateCoinCode(trimmedCoinCode)
		country.CoinCode = trimmedCoinCode
	}
	if params.CoinSymbol != nil {
		trimmedCoinSymbol := strings.TrimSpace(*params.CoinSymbol)
		validator.ValidateCoinSymbol(trimmedCoinSymbol)
		country.CoinSymbol = trimmedCoinSymbol
	}

	if validator.HasErrors() {
		return nil, apperror.NewValidationError("CountryService", validator.Errors())
	}

	country.Touch()

	if err := s.countryRepository.Update(ctx, country); err != nil {
		return nil, err
	}

	return country, nil
}

func (s *Service) DeleteCountry(ctx context.Context, id uuid.UUID) error {
	if _, err := s.countryRepository.GetByID(ctx, id); err != nil {
		return err
	}
	return s.countryRepository.Delete(ctx, id)
}
