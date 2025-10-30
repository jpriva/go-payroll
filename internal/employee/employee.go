package employee

import (
	"context"
	"payroll/internal/domain"

	"payroll/internal/apperror"
	"strings"
	"time"

	"github.com/google/uuid"
)

const employeeOrigin = "Employee"

type EmployeeGender string

const (
	GenderMale   EmployeeGender = "MALE"
	GenderFemale EmployeeGender = "FEMALE"
	GenderOther  EmployeeGender = "OTHER"
)

func (g EmployeeGender) IsValid() bool {
	switch g {
	case GenderMale, GenderFemale, GenderOther:
		return true
	}
	return false
}

type Employee struct {
	domain.BaseEntity
	TenantID    uuid.UUID
	WorkspaceID uuid.UUID

	FirstName string
	LastName  string
	Email     string
	Address   string
	DocTypeID uuid.UUID
	DocNumber string
	//Not obligatory
	BirthDate *time.Time
	Gender    *EmployeeGender
	Phone     *string
}

type CreateEmployeeParams struct {
	TenantID    uuid.UUID
	WorkspaceID uuid.UUID
	FirstName   string
	LastName    string
	Email       string
	Address     string
	DocTypeID   uuid.UUID
	DocNumber   string
	BirthDate   *time.Time
	Gender      *string
	Phone       *string
}

type UpdateEmployeeParams struct {
	FirstName *string
	LastName  *string
	Email     *string
	Address   *string
	DocTypeID *uuid.UUID
	DocNumber *string
	BirthDate *time.Time
	Gender    *string
	Phone     *string
}

func NewEmployee(params CreateEmployeeParams) (*Employee, error) {
	validator := NewValidator()

	if params.TenantID == uuid.Nil {
		validator.errs["TenantID"] = "is empty"
	}
	if params.WorkspaceID == uuid.Nil {
		validator.errs["WorkspaceID"] = "is empty"
	}

	params.FirstName = strings.TrimSpace(params.FirstName)
	params.LastName = strings.TrimSpace(params.LastName)
	params.Email = strings.TrimSpace(params.Email)
	params.Address = strings.TrimSpace(params.Address)
	params.DocNumber = strings.TrimSpace(params.DocNumber)
	if params.Phone != nil {
		*params.Phone = strings.TrimSpace(*params.Phone)
	}

	validator.ValidateFirstName(params.FirstName)
	validator.ValidateLastName(params.LastName)
	validator.ValidateEmail(params.Email)
	validator.ValidateBirthDate(params.BirthDate)
	validator.ValidateDocTypeID(params.DocTypeID)
	validator.ValidateDocNumber(params.DocNumber)
	validator.ValidateGender(params.Gender)
	validator.ValidatePhone(params.Phone)

	var empGender *EmployeeGender
	if params.Gender != nil {
		gender := EmployeeGender(*params.Gender)
		empGender = &gender
	}
	if validator.HasErrors() {
		return nil, apperror.NewValidationError(employeeOrigin, validator.Errors())
	}

	emp := &Employee{
		TenantID:    params.TenantID,
		WorkspaceID: params.WorkspaceID,
		FirstName:   params.FirstName,
		LastName:    params.LastName,
		Email:       params.Email,
		Address:     params.Address,
		DocTypeID:   params.DocTypeID,
		DocNumber:   params.DocNumber,
		BirthDate:   params.BirthDate,
		Gender:      empGender,
		Phone:       params.Phone,
	}
	emp.Initialize()

	return emp, nil
}

type Repository interface {
	Create(ctx context.Context, employee *Employee) error
	ListByWorkspaceIDAndTenantID(ctx context.Context, workspaceID uuid.UUID, tenantID uuid.UUID) ([]*Employee, error)
	GetByID(ctx context.Context, id uuid.UUID) (*Employee, error)
	Update(ctx context.Context, employee *Employee) error
	Delete(ctx context.Context, id uuid.UUID) error
	ExistsByTenantIDAndDocNumber(ctx context.Context, tenantID uuid.UUID, docNumber string) (bool, error)
	ExistsByTenantIDAndEmail(ctx context.Context, tenantID uuid.UUID, email string) (bool, error)
}
