package domain

import (
	"errors"
	"sort"
	"time"

	"github.com/beevik/guid"
)

//Product entity
type Product struct {
	ID          string     `bson:"_id" json:"id"`
	Name        string     `bson:"name" json:"name"`
	Description string     `bson:"description" json:"description"`
	Skus        []Sku      `bson:"skus" json:"skus"`
	Active      bool       `bson:"active" json:"active"`
	Categories  []Category `bson:"categories" json:"categories"`
	CreateAt    time.Time  `bson:"create_at" json:"create_at"`
	UpdateAt    time.Time  `bson:"update_at" json:"update_at"`
}

//NewProduct constructor
func NewProduct(
	name string,
	descripion string,
	categories []Category,
) (Product, error) {
	if len(categories) == 0 {
		return Product{}, errors.New("can't find product category")
	}

	newguid := guid.New()

	return Product{
		ID:          newguid.String(),
		Name:        name,
		Description: descripion,
		Active:      true,
		UpdateAt:    DateNow(),
		CreateAt:    DateNow(),
		Skus:        []Sku{},
		Categories:  categories,
	}, nil
}

//SetName update product
func (p *Product) SetName(name string) {
	p.Name = name
	p.UpdateAt = DateNow()

}

//SetDescription update product
func (p *Product) SetDescription(description string) {
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

//HasVariations return if product has variations
func (p *Product) HasVariations() bool {
	return len(p.Skus) > 0
}

//AddSku add product variation
func (p *Product) AddSku(sku Sku) {
	p.Skus = append(p.Skus, sku)
	p.UpdateAt = DateNow()
}

//RemoveSku remove product variation
func (p *Product) RemoveSku(skuID string) error {
	index := sort.Search(len(p.Skus), func(i int) bool {
		return p.Skus[i].ID == skuID
	})

	if len(p.Skus) == index {
		return errors.New("sku not found")
	}

	p.Skus = append(p.Skus[:index], p.Skus[index+1:]...)
	p.UpdateAt = DateNow()
	return nil
}

//AddCategory add product category
func (p *Product) AddCategory(category Category) {
	p.Categories = append(p.Categories, category)
	p.UpdateAt = DateNow()
}

//RemoveCategory remove product category
func (p *Product) RemoveCategory(categoryID string) error {
	index := sort.Search(len(p.Categories), func(i int) bool {
		return p.Categories[i].ID == categoryID
	})

	if len(p.Categories) == index {
		return errors.New("category not found")
	}

	p.Categories = append(p.Categories[:index], p.Categories[index+1:]...)
	p.UpdateAt = DateNow()
	return nil
}
