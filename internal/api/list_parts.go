package api

import (
	"context"
	"inventory/internal/api/converter"
	"inventory/internal/entity"
	in "inventory/pkg/pb"
)

func (s *InventoryServer) ListParts(_ context.Context, req *in.ListPartsRequest) (*in.ListPartsResponse, error) {
	parts, err := s.Usecase.ListParts(converter.ProtoFilterToEntity(req.Filter))
	if err != nil {
		return nil, entity.ErrDataBaseFailed
	}

	convertedParts := converter.PartsEntityToProto(parts)

	return &in.ListPartsResponse{Parts: convertedParts}, nil
}
