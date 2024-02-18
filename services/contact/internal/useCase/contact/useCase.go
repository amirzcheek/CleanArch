package contact

import "architecture_go/services/contact/internal/useCase/adapters/storage"

type UseCase struct {
	adapterStorage storage.Contact
}

func New(storage storage.Contact) *UseCase {
	var uc = &UseCase{
		adapterStorage: storage,
	}
	return uc
}
