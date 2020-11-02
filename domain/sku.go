package domain

import "github.com/beevik/guid"

//Sku product variation
type Sku struct {
	Name     string  `bson:"name" json:"name"`
	Price    float64 `bson:"price" json:"price"`
	Quantity int     `bson:"quantity" json:"quantity"`
	Aggreate
}

//NewSku constructor
func NewSku(name string, price float64, quantity int) Sku {
	sku := Sku{
		Aggreate: Aggreate{
			ID: guid.NewString(),
		},
		Name:     name,
		Quantity: quantity,
		Price:    price,
	}

	return sku
}
