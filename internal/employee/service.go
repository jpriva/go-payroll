package employee

import (
	"context"
	"payroll/internal/apperror"
	"payroll/internal/doctype"
	"payroll/internal/platform/logger"
	"payroll/internal/workspace"
	"strings"

	"github.com/google/uuid"
)

type Service struct {
	employeeRepo  Repository
	workspaceRepo workspace.Repository
	docTypeRepo   doctype.Repository
	logger        logger.Logger
}

func NewService(er Repository, wr workspace.Repository, dtr doctype.Repository, l logger.Logger) *Service {
	return &Service{
		employeeRepo:  er,
		workspaceRepo: wr,
		docTypeRepo:   dtr,
		logger:        l,
	}
}

func (s *Service) Create(ctx context.Context, params CreateEmployeeParams) (*Employee, error) {
	employee, err := NewEmployee(params)
	if err != nil {
		s.logger.Warn("Failed to create new employee due to validation errors", "errors", err)
		return nil, err
	}

	ws, err := s.workspaceRepo.Get(ctx, params.WorkspaceID)
	if err != nil {
		return nil, apperror.New(apperror.TypeInvalid, "EmployeeService", "Invalid WorkspaceID")
	}
	s.logger.Debug("Workspace validation successful", "workspace_id", params.WorkspaceID)

	isValid, err := s.docTypeRepo.IsValidForCountry(ctx, params.DocTypeID, ws.CountryID)
	if err != nil {
		s.logger.Error(err, "Failed to validate document type for country")
		return nil, err
	}
	if !isValid {
		err := apperror.New(apperror.TypeInvalid, "EmployeeService", "DocType is not valid for the employee's country")
		s.logger.Warn(err.Error(), "doc_type_id", params.DocTypeID, "country", ws.CountryID)
		return nil, err
	}

	if err := s.employeeRepo.Create(ctx, employee); err != nil {
		s.logger.Error(err, "Failed to save employee to repository")
		return nil, err
	}

	s.logger.Info("Employee created successfully", "employee_id", employee.ID)
	return employee, nil
}

func (s *Service) GetByID(ctx context.Context, id uuid.UUID) (*Employee, error) {
	return s.employeeRepo.GetByID(ctx, id)
}

func (s *Service) Update(ctx context.Context, id uuid.UUID, params UpdateEmployeeParams) (*Employee, error) {
	employee, err := s.employeeRepo.GetByID(ctx, id)
	if err != nil {
		s.logger.Error(err, "Failed to get employee for update", "employee_id", id)
		return nil, err
	}

	validator := NewValidator()

	if params.FirstName != nil {
		*params.FirstName = strings.TrimSpace(*params.FirstName)
		validator.ValidateFirstName(*params.FirstName)
		employee.FirstName = *params.FirstName
	}
	if params.LastName != nil {
		*params.LastName = strings.TrimSpace(*params.LastName)
		validator.ValidateLastName(*params.LastName)
		employee.LastName = *params.LastName
	}
	if params.Email != nil {
		*params.Email = strings.TrimSpace(*params.Email)
		validator.ValidateEmail(*params.Email)
		employee.Email = *params.Email
	}
	if params.DocTypeID != nil {
		validator.ValidateDocTypeID(*params.DocTypeID)
		employee.DocTypeID = *params.DocTypeID
	}
	if params.DocNumber != nil {
		*params.DocNumber = strings.TrimSpace(*params.DocNumber)
		validator.ValidateDocNumber(*params.DocNumber)
		employee.DocNumber = *params.DocNumber
	}
	if params.BirthDate != nil {
		if params.BirthDate.IsZero() {
			employee.BirthDate = nil
		} else {
			validator.ValidateBirthDate(params.BirthDate)
			employee.BirthDate = params.BirthDate
		}
	}
	if params.Gender != nil {
		*params.Gender = strings.TrimSpace(*params.Gender)
		if *params.Gender == "" {
			employee.Gender = nil
		} else {
			validator.ValidateGender(params.Gender)
			gender := EmployeeGender(*params.Gender)
			employee.Gender = &gender
		}
	}
	if params.Phone != nil {
		*params.Phone = strings.TrimSpace(*params.Phone)
		if *params.Phone == "" {
			employee.Phone = nil
		} else {
			validator.ValidatePhone(params.Phone)
			employee.Phone = params.Phone
		}
	}

	if validator.HasErrors() {
		err := apperror.NewValidationError("UpdateEmployee", validator.Errors())
		s.logger.Warn("Failed to update employee due to validation errors", "errors", err)
		return nil, err
	}

	employee.Touch()

	if err := s.employeeRepo.Update(ctx, employee); err != nil {
		s.logger.Error(err, "Failed to save updated employee to repository", "employee_id", id)
		return nil, err
	}
	s.logger.Info("Employee updated successfully", "employee_id", employee.ID)
	return employee, nil
}

func (s *Service) ListByWorkspaceIDAndTenantID(ctx context.Context, workspaceID uuid.UUID, tenantID uuid.UUID) ([]*Employee, error) {
	return s.employeeRepo.ListByWorkspaceIDAndTenantID(ctx, workspaceID, tenantID)
}

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	if _, err := s.employeeRepo.GetByID(ctx, id); err != nil {
		return err
	}
	return s.employeeRepo.Delete(ctx, id)
}
