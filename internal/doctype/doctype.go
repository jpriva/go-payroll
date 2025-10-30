package doctype

import (
	"context"

	"github.com/google/uuid"
)

type DocType struct {
	ID        uuid.UUID
	CountryId uuid.UUID
	Code      string
	Name      string
}

type Repository interface {
	IsValidForCountry(ctx context.Context, docTypeID uuid.UUID, countryID uuid.UUID) (bool, error)
	Get(ctx context.Context, id uuid.UUID) (*DocType, error)
}
