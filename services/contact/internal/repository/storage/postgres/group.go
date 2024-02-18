package postgres

import (
	"architecture_go/services/contact/internal/domain/group"
	"github.com/google/uuid"
)

func CreateGroup(group *group.Group) (*group.Group, error) {
	return nil, nil
}
func UpdateGroup(ID uuid.UUID, groupUpdate *group.Group) (*group.Group, error) {
	return nil, nil
}

func ListGroups() ([]*group.Group, error) {
	return nil, nil
}

func ReadGroupByID(ID uuid.UUID) (*group.Group, error) {
	return nil, nil
}
