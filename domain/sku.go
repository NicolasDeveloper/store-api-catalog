package domain

import "github.com/beevik/guid"

//Sku product variation
type Sku struct {
	ID       string  `bson:"id" json:"id"`
	Name     string  `bson:"name" json:"name"`
	Price    float64 `bson:"price" json:"price"`
	Quantity int     `bson:"quantity" json:"quantity"`
}

//NewSku constructor
func NewSku(name string, price string, quantity int) Sku {
	newguid := guid.New()

	return Sku{
		ID:       newguid.String(),
		Name:     name,
		Quantity: quantity,
	}
}
