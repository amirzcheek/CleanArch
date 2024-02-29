package group

import (
	"architecture_go/services/contact/internal/domain/group"
	"context"
	"github.com/google/uuid"
)

func (uc *UseCase) Create(ctx context.Context, groupCreate *group.Group) (*group.Group, error) {
	createdGroup, err := uc.adapterStorage.CreateGroup(ctx, groupCreate)
	if err != nil {
		return nil, err
	}
	return createdGroup, nil
}

func (uc *UseCase) Update(ctx context.Context, ID uuid.UUID, groupUpdate *group.Group) (*group.Group, error) {
	updatedGroup, err := uc.adapterStorage.UpdateGroup(ctx, ID, groupUpdate)
	if err != nil {
		return nil, err
	}
	return updatedGroup, nil
}

func (uc *UseCase) Delete(ctx context.Context, ID uuid.UUID) error {
	err := uc.adapterStorage.DeleteGroup(ctx, ID)
	if err != nil {
		return err
	}
	return nil
}

func (uc *UseCase) List(ctx context.Context) ([]*group.Group, error) {
	groups, err := uc.adapterStorage.ListGroups(ctx)
	if err != nil {
		return nil, err
	}
	return groups, nil
}

func (uc *UseCase) ReadByID(ctx context.Context, ID uuid.UUID) (*group.Group, error) {
	readGroup, err := uc.adapterStorage.ReadGroupByID(ctx, ID)
	if err != nil {
		return nil, err
	}
	return readGroup, nil
}
