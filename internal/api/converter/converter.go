package converter

import (
	"inventory/internal/entity"
	in "inventory/pkg/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func PartEntityToProto(part *entity.Part) in.Part {
	return in.Part{
		Name:          part.Name,
		UUID:          part.UUID,
		Tags:          part.Tags,
		Price:         part.Price,
		Description:   part.Description,
		StockQuantity: part.StockQuantity,
		CreatedAt:     timestamppb.New(part.CreatedAt),
		UpdatedAt:     timestamppb.New(part.UpdatedAt),
		Category:      CategoryToProto(part.Category),
		Dimensions:    DimensionsToProto(part.Dimensions),
		Manufacturer:  ManufacturerToProto(part.Manufacturer),
	}
}
func PartsEntityToProto(parts []*entity.Part) []*in.Part {
	convertedParts := make([]*in.Part, len(parts))

	for _, part := range parts {
		convertedPart := PartEntityToProto(part)
		convertedParts = append(convertedParts, &convertedPart)
	}
	return convertedParts
}

func ManufacturerToProto(manu entity.Manufacturer) *in.Manufacturer {
	return &in.Manufacturer{
		Name:    manu.Name,
		Country: manu.Country,
		Website: manu.Website,
	}
}

func DimensionsToProto(dims entity.Dimensions) *in.Dimensions {
	return &in.Dimensions{
		Width:  dims.Width,
		Height: dims.Height,
		Length: dims.Length,
		Weight: dims.Weight,
	}
}

func CategoryToProto(cat entity.Category) in.Category {
	switch cat {
	case entity.CategoryUnknown:
		return in.Category_CATEGORY_ENGINE
	case entity.CategoryFuel:
		return in.Category_CATEGORY_FUEL

	case entity.CategoryWing:
		return in.Category_CATEGORY_WING
	case entity.CategoryPorthole:
		return in.Category_CATEGORY_PORTHOLE
	}
	return in.Category_CATEGORY_UNKNOWN
}

func ProtoCategoryToEntity(cat []in.Category) []entity.Category {
	categories := make([]entity.Category, len(cat))
	for _, cate := range cat {
		switch cate {
		case in.Category_CATEGORY_WING:
			categories = append(categories, entity.CategoryWing)
		case in.Category_CATEGORY_PORTHOLE:
			categories = append(categories, entity.CategoryPorthole)
		case in.Category_CATEGORY_FUEL:
			categories = append(categories, entity.CategoryFuel)
		case in.Category_CATEGORY_ENGINE:
			categories = append(categories, entity.CategoryEngine)
		default:
			categories = append(categories, entity.CategoryUnknown)
		}
	}
	return categories
}

func ProtoFilterToEntity(filter *in.PartsFilter) entity.PartsFilter {
	return entity.PartsFilter{
		ManufacturerContries: filter.ManufacturerCountries,
		Names:                filter.Names,
		UUIDS:                filter.Uuids,
		Categories:           ProtoCategoryToEntity(filter.Categories),
		Tags:                 filter.Tags,
	}
}
