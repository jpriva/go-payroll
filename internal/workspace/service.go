package workspace

import (
	"context"
	"payroll/internal/apperror"

	"github.com/google/uuid"
)

const serviceOrigin = "Workspace Service"

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) Create(ctx context.Context, params CreateWorkspaceParams) (*Workspace, error) {
	workspace, err := NewWorkspace(params)
	if err != nil {
		return nil, err
	}

	exists, err := s.repo.ExistsByTenantIDAndCode(ctx, workspace.TenantID, workspace.Code)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, apperror.New(apperror.TypeDuplicate, serviceOrigin,
			"a workspace with this code already exists for the given tenant")
	}

	if err := s.repo.Create(ctx, workspace); err != nil {
		return nil, err
	}

	return workspace, nil
}

func (s *Service) Get(ctx context.Context, id uuid.UUID) (*Workspace, error) {
	return s.repo.Get(ctx, id)
}

func (s *Service) Update(ctx context.Context, id uuid.UUID, params UpdateWorkspaceParams) (*Workspace, error) {
	ws, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	validator := NewValidator()

	if params.Code != nil {
		validator.ValidateCode(*params.Code)
		ws.Code = *params.Code
	}
	if params.Name != nil {
		validator.ValidateName(*params.Name)
		ws.Name = *params.Name
	}
	if params.Status != nil {
		validator.ValidateStatus(params.Status)
		ws.Status = *params.Status
	}

	if validator.HasErrors() {
		return nil, apperror.NewValidationError(serviceOrigin, validator.Errors())
	}

	ws.Touch()

	if err := s.repo.Update(ctx, ws); err != nil {
		return nil, err
	}

	return ws, nil
}

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	if _, err := s.repo.Get(ctx, id); err != nil {
		return err
	}
	return s.repo.Delete(ctx, id)
}
