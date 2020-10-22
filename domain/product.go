package domain

import (
	"time"

	"github.com/beevik/guid"
)

//Product entity
type Product struct {
	ID          string    `bson:"_id" json:"id"`
	Name        string    `bson:"name" json:"name"`
	Price       float64   `bson:"price" json:"price"`
	Description string    `bson:"description" json:"description"`
	Active      bool      `bson:"active" json:"active"`
	CreateAt    time.Time `bson:"create_at" json:"create_at"`
	UpdateAt    time.Time `bson:"update_at" json:"update_at"`
}

//NewProduct constructor
func NewProduct(
	name string,
	price float64,
	descripion string,
) (Product, error) {
	newguid := guid.New()

	return Product{
		ID:          newguid.String(),
		Name:        name,
		Price:       price,
		Description: descripion,
		Active:      true,
		UpdateAt:    time.Now(),
		CreateAt:    time.Now(),
	}, nil
}

//Update update product
func (p *Product) Update(name string, price float64, description string) {
	p.Name = name
	p.Price = price
	p.Description = description
	p.UpdateAt = time.Now()
}

//Disable inactive product to not be used anymore
func (p *Product) Disable() {
	p.Active = false
	p.UpdateAt = time.Now()
}

//Enable active product to not be used
func (p *Product) Enable() {
	p.Active = true
	p.UpdateAt = time.Now()
}
