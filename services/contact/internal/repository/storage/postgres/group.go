package postgres

import (
	"architecture_go/services/contact/internal/domain/group"
	"context"
	"github.com/google/uuid"
)

func (r *Repository) CreateGroup(ctx context.Context, groupCreate *group.Group) (*group.Group, error) {
	return nil, nil
}
func (r *Repository) UpdateGroup(ctx context.Context, ID uuid.UUID, groupUpdate *group.Group) (*group.Group, error) {
	return nil, nil
}

func (r *Repository) ListGroups(ctx context.Context) ([]*group.Group, error) {
	return nil, nil
}
func (r *Repository) DeleteGroup(ctx context.Context, ID uuid.UUID) error {
	return nil
}

func (r *Repository) ReadGroupByID(ctx context.Context, ID uuid.UUID) (*group.Group, error) {
	return nil, nil
}
