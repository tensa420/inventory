package api

import (
	"context"
	"inventory/internal/api/converter"
	"inventory/internal/entity"
	"time"

	in "inventory/pkg/pb"
)

func (s *InventoryServer) GetPart(_ context.Context, req *in.GetPartRequest) (*in.GetPartResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	part, err := s.Usecase.GetPart(ctx, req.Uuid)
	if err != nil {
		return nil, entity.ErrNotFound
	}
	convertedPart := converter.PartEntityToProto(part)

	return &in.GetPartResponse{Part: &convertedPart}, nil
}
