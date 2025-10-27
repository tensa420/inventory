package api

import (
	"inventory/internal/usecase"
	v1 "inventory/pkg/pb"
)

type InventoryServer struct {
	Usecase *usecase.InventoryUsecase
	v1.UnimplementedInventoryServiceServer
}

func NewInventoryServer(usecase *usecase.InventoryUsecase) *InventoryServer {
	return &InventoryServer{
		Usecase: usecase,
	}
}
