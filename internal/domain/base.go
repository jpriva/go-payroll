package domain

import (
	"time"

	"github.com/google/uuid"
)

type BaseEntity struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (b *BaseEntity) Initialize() {
	b.ID = uuid.Must(uuid.NewV7())
	now := time.Now().UTC()
	b.CreatedAt = now
	b.UpdatedAt = now
}

func (b *BaseEntity) Touch() {
	b.UpdatedAt = time.Now().UTC()
}
