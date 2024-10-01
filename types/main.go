package types

type CpArticleAttributes struct {
	ID                 string      `json:"id"`
	SKU                string      `json:"sku"`
	Name               string      `json:"name"`
	DNE                string      `json:"dne"`
	Type               string      `json:"type"`
	Image              string      `json:"image"`
	Link               string      `json:"link"`
	Price              float64     `json:"price"`
	TotalPrice         float64     `json:"total_price"`
	Discount           float64     `json:"discount"`
	Stock              int         `json:"stock"`
	Rate               float64     `json:"rate"`
	Review             int         `json:"review"`
	Featured           bool        `json:"featured"`
	New                bool        `json:"new"`
	KickbackSaf        bool        `json:"kickback_saf"`
	CategoryID         string      `json:"category_id"`
	Manufacturer       string      `json:"manufacturer"`
	ManufacturerID     string      `json:"manufacturer_id"`
	Multiple           bool        `json:"multiple"`
	Category           string      `json:"category"`
	Quantity           int         `json:"quantity"`
	UnitPrice          float64     `json:"unit_price"`
	Relevance          interface{} `json:"relevance"`
	Attributes         []Attribute `json:"attributes"`
	Specials           Specials    `json:"specials"`
	Alerts             Alerts      `json:"alerts"`
	ConfigurationPrice float64     `json:"configuration_price"`
	TypeStorage        string      `json:"type_storage"`
}

type Attribute struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Value string `json:"value"`
}

type Specials struct {
	Capacity            string `json:"capacity"`
	Modules             int    `json:"modules"`
	DVI                 int    `json:"dvi"`
	HDMI                int    `json:"hdmi"`
	RamSlots            string `json:"ramSlots"`
	MaxMemory           string `json:"maxMemory"`
	DisplayPort         int    `json:"displayPort"`
	FansAvailable       int    `json:"fansAvailable"`
	FansInstalled       int    `json:"fansInstalled"`
	IncludedPowerSupply string `json:"includedPowerSupply"`
	GraphicIncluded     string `json:"graphicIncluded"`
}

type Alerts struct {
	PriceDown bool `json:"priceDown"`
	PriceUp   bool `json:"priceUp"`
	LowStock  bool `json:"lowStock"`
}

type CpArticleComponent struct {
	Type          string              `json:"type"`
	ID            string              `json:"id"`
	Quantity      int                 `json:"quantity"`
	Attributes    CpArticleAttributes `json:"attributes"`
	Relationships []Relationship      `json:"relationships"`
	Included      []IncludedComponent `json:"included"`
}

type Relationship struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type IncludedComponent struct {
	Type       string             `json:"type"`
	ID         string             `json:"id"`
	Attributes IncludedAttributes `json:"attributes"`
	Links      IncludedLinks      `json:"links"`
}

type IncludedAttributes struct {
	Title string `json:"title"`
}

type IncludedLinks struct {
	Self string `json:"self"`
}

type FilterSectionItem struct {
	GroupID  string `json:"groupId"`
	Title    string `json:"title"`
	Value    string `json:"value"`
	Selected bool   `json:"selected"`
	Count    int    `json:"count"`
	Category string `json:"category"`
}

type FilterSectionRow struct {
	ID         string              `json:"id"`
	Title      string              `json:"title"`
	CheckBoxes []FilterSectionItem `json:"checkBoxes"`
	Max        int                 `json:"max"`
}

type IFiltersApplied struct {
	Category string              `json:"category"`
	Value    string              `json:"value"`
	Min      float64             `json:"min"`
	Max      float64             `json:"max"`
	Filters  []FilterSectionItem `json:"filters"`
}

type IFilters struct {
	Data             []CpArticleComponent `json:"data"`
	OnlyStock        bool                 `json:"onlyStock"`
	PaginateFilters  []FilterSectionRow   `json:"paginateFilters"`
	ComponentFilters []ComponentFilter    `json:"componentFilters"`
	FiltersApplied   []IFiltersApplied    `json:"filtersApplied"`
	Search           string               `json:"search"`
	Price            PriceRange           `json:"price"`
	StorageType      []string             `json:"storageType"`
}

type ComponentFilter struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type PriceRange struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}

type FiltersToSearch struct {
	GroupID string
	Filters []string
}

// type Grouped struct {
// 	Key     string
// 	Filters FilterSectionItem
// }
