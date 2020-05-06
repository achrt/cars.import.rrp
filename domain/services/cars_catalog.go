package services

type Price struct {
	Id            string              `json:"id"`
	Type          string              `json:"type"`
	Attributes    *PriceAttributes    `json:"attributes"`
	Relationships *PriceRelationships `json:"relationships"`
}

type PriceAttributes struct {
	Values []*PriceValue `json:"values"`
	Status uint8         `json:"status"`
}

type PriceValue struct {
	PriceTypeId string `json:"price_type_id"`
	Value       uint64 `json:"value"`
}

type PriceRelationships struct {
	MarketingComplectation *PriceMarketingComplectationRelationship `json:"marketing_complectation"`
	PriceType              *MultipleRelationWithName                `json:"price_type"`
}

type PriceMarketingComplectationRelationship struct {
	Data *Relation `json:"data"`
}

type PriceType struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Code string `json:"code,omitempty"`
}

type MultipleRelationWithName struct {
	Data []*RelationWithName `json:"data"`
}

type RelationWithName struct {
	Id   string `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
	Code string `json:"code,omitempty"`
}

type Relation struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type Services interface {
	CarsCatalog() CarsCatalog
}

type CarsCatalog interface {
	GetPricesByMarkId(markId string) (*Price, error)
}
