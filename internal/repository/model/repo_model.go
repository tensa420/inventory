package model

import (
	"time"

	"inventory/internal/entity"
)

type Part struct {
	UUID          string
	Name          string
	Description   string
	Price         float64
	StockQuantity int64
	Category      Category
	Dimensions    Dimensions
	Manufacturer  Manufacturer
	Tags          []string
	Metadata      map[string]Value
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
type Dimensions struct {
	length float64
	width  float64
	height float64
	weight float64
}
type Category string

const (
	CategoryUnknown  Category = "Unknown"
	CategoryEngine   Category = "Engine"
	CategoryFuel     Category = "Fuel"
	CategoryPorthole Category = "Portholle"
	CategoryWing     Category = "Wing"
)

type Manufacturer struct {
	name    string
	country string
	website string
}

type Value struct {
	string_value string
	int64_value  int64
	double_value float64
	bool_value   bool
}

func ConvertRepoModelToEntity(part Part) *entity.Part {
	return &entity.Part{
		UUID:          part.UUID,
		Name:          part.Name,
		Description:   part.Description,
		Price:         part.Price,
		StockQuantity: part.StockQuantity,
		Category:      entity.Category(part.Category),
		Dimensions:    RepoDimensionsToEntity(part.Dimensions),
		Manufacturer:  RepoManufacturerToEntity(part.Manufacturer),
		Tags:          part.Tags,
		CreatedAt:     part.CreatedAt,
		UpdatedAt:     part.UpdatedAt,
	}
}

func RepoDimensionsToEntity(dime Dimensions) entity.Dimensions {
	return entity.Dimensions{
		Length: dime.length,
		Width:  dime.width,
		Height: dime.height,
		Weight: dime.weight,
	}
}
func RepoManufacturerToEntity(part Manufacturer) entity.Manufacturer {
	return entity.Manufacturer{
		Name:    part.name,
		Country: part.country,
		Website: part.website,
	}
}
