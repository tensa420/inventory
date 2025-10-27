package usecase

import (
	"context"

	"inventory/internal/entity"
)

func (u *InventoryUsecase) GetPart(ctx context.Context, partUUID string) (*entity.Part, error) {

	part, err := u.repo.GetPart(partUUID)
	if err != nil {
		return nil, entity.ErrNotFound
	}

	return part, nil
}
