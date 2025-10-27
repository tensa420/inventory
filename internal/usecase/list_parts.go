package usecase

import "inventory/internal/entity"

func (u *InventoryUsecase) ListParts(filter entity.PartsFilter) ([]*entity.Part, error) {
	parts, err := u.repo.ListParts(filter)
	if err != nil {
		return nil, entity.ErrDataBaseFailed
	}
	return parts, nil
}
