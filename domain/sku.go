package domain

import "github.com/beevik/guid"

//Sku product variation
type Sku struct {
	Name     string  `bson:"name" json:"name"`
	Price    float64 `bson:"price" json:"price"`
	Quantity int     `bson:"quantity" json:"quantity"`
	Entity
}

//NewSku constructor
func NewSku(name string, price string, quantity int) Sku {
	sku := Sku{
		Entity: Entity{
			ID: guid.NewString(),
		},
		Name:     name,
		Quantity: quantity,
	}

	return sku
}
