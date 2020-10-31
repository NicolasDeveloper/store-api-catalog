package domain

import (
	"errors"
	"time"

	"github.com/beevik/guid"
)

//Product entity
type Product struct {
	ID          string    `bson:"_id" json:"id"`
	Name        string    `bson:"name" json:"name"`
	Description string    `bson:"description" json:"description"`
	Skus        []Sku     `bson:"skus" json:"skus"`
	Active      bool      `bson:"active" json:"active"`
	Categories  []string  `bson:"categories" json:"categories"`
	CreateAt    time.Time `bson:"create_at" json:"create_at"`
	UpdateAt    time.Time `bson:"update_at" json:"update_at"`
}

//NewProduct constructor
func NewProduct(
	name string,
	descripion string,
) (Product, error) {
	newguid := guid.New()

	return Product{
		ID:          newguid.String(),
		Name:        name,
		Description: descripion,
		Active:      true,
		UpdateAt:    DateNow(),
		CreateAt:    DateNow(),
		Skus:        []Sku{},
		Categories:  []string{},
	}, nil
}

//Update update product
func (p *Product) Update(name string, description string) {
	p.Name = name
	p.Description = description
	p.UpdateAt = DateNow()
}

//Disable inactive product to not be used anymore
func (p *Product) Disable() {
	p.Active = false
	p.UpdateAt = DateNow()
}

//Enable active product to not be used
func (p *Product) Enable() {
	p.Active = true
	p.UpdateAt = DateNow()
}

//AddSku add product variation
func (p *Product) AddSku(sku Sku) {
	p.Skus = append(p.Skus, sku)
	p.UpdateAt = DateNow()
}

//AddCategory add product category
func (p *Product) AddCategory(categoryID string) error {
	if guid.IsGuid(categoryID) == false {
		return errors.New("Cannot add this category because its not identification")
	}

	p.Categories = append(p.Categories, categoryID)
	p.UpdateAt = DateNow()
	return nil
}
