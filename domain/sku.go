package domain

//Sku product variation
type Sku struct {
	Name     string  `bson:"name" json:"name"`
	Price    float64 `bson:"price" json:"price"`
	Quantity float64 `bson:"quantity" json:"quantity"`
}
