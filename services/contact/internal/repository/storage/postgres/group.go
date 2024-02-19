package postgres

import (
	"architecture_go/services/contact/internal/domain/group"
	"github.com/google/uuid"
)

func (r *Repository) CreateGroup(groupCreate *group.Group) (*group.Group, error) {
	return nil, nil
}
func (r *Repository) UpdateGroup(ID uuid.UUID, groupUpdate *group.Group) (*group.Group, error) {
	return nil, nil
}

func (r *Repository) ListGroups() ([]*group.Group, error) {
	return nil, nil
}
func (r *Repository) DeleteGroup(ID uuid.UUID) error {
	return nil
}

func (r *Repository) ReadGroupByID(ID uuid.UUID) (*group.Group, error) {
	return nil, nil
}
