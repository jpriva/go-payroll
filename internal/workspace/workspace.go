package workspace

import (
	"context"
	"payroll/internal/apperror"
	"payroll/internal/domain"

	"github.com/google/uuid"
)

type WorkspaceStatus string

const modelOrigin = "Workspace"

const (
	WorkspaceStatusActive   WorkspaceStatus = "ACTIVE"
	WorkspaceStatusInactive WorkspaceStatus = "INACTIVE"
	WorkspaceStatusPending  WorkspaceStatus = "PENDING"
)

func (s WorkspaceStatus) IsValid() bool {
	switch s {
	case WorkspaceStatusActive, WorkspaceStatusInactive, WorkspaceStatusPending:
		return true
	}
	return false
}

type Workspace struct {
	domain.BaseEntity
	TenantID  uuid.UUID
	Code      string
	Name      string
	Status    WorkspaceStatus
	CountryID uuid.UUID
}

type CreateWorkspaceParams struct {
	TenantID  uuid.UUID
	CountryID uuid.UUID
	Code      string
	Name      string
	Status    *WorkspaceStatus
}

type UpdateWorkspaceParams struct {
	Code   *string
	Name   *string
	Status *WorkspaceStatus
}

func NewWorkspace(params CreateWorkspaceParams) (*Workspace, error) {
	validator := NewValidator()

	if params.TenantID == uuid.Nil {
		validator.AddError("TenantID", "is empty")
	}

	validator.ValidateCode(params.Code)
	validator.ValidateName(params.Name)
	validator.ValidateCountryID(params.CountryID)
	validator.ValidateStatus(params.Status)

	var status WorkspaceStatus
	if params.Status == nil {
		status = WorkspaceStatusPending
	} else {
		status = *params.Status
	}

	if validator.HasErrors() {
		return nil, apperror.NewValidationError(modelOrigin, validator.Errors())
	}

	ws := &Workspace{
		TenantID:  params.TenantID,
		Code:      params.Code,
		Name:      params.Name,
		CountryID: params.CountryID,
		Status:    status,
	}
	ws.Initialize()

	return ws, nil
}

type Repository interface {
	Create(ctx context.Context, ws *Workspace) error
	Get(ctx context.Context, id uuid.UUID) (*Workspace, error)
	Update(ctx context.Context, ws *Workspace) error
	Delete(ctx context.Context, id uuid.UUID) error
	ExistsByTenantIDAndCode(ctx context.Context, tenantID uuid.UUID, code string) (bool, error)
}
