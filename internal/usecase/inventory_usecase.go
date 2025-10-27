package usecase

import "inventory/internal/repository/repository"

type InventoryUsecase struct {
	repo repository.InventoryRepository
}

func NewInventoryUsecase(repo *repository.InventoryRepository) *InventoryUsecase {
	return &InventoryUsecase{
		repo: *repo,
	}
}
