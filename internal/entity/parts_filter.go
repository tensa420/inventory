package entity

type PartsFilter struct {
	UUIDS                []string
	Names                []string
	Categories           []Category
	ManufacturerContries []string
	Tags                 []string
}
