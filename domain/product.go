package domain

import (
	"errors"
	"sort"
	"time"

	"github.com/beevik/guid"
)

//Product entity
type Product struct {
	Name          string    `bson:"name" json:"name"`
	Description   string    `bson:"description" json:"description"`
	Active        bool      `bson:"active" json:"active"`
	Skus          []Sku     `bson:"skus" json:"skus"`
	CategoriesIds []string  `bson:"categories" json:"categories"`
	CreateAt      time.Time `bson:"create_at" json:"create_at"`
	UpdateAt      time.Time `bson:"update_at" json:"update_at"`
	AggreateRoot  `bson:",inline"`
}

//NewProduct constructor
func NewProduct(
	name string,
	descripion string,
	categoriesIDs []string,
) (Product, error) {
	if len(categoriesIDs) == 0 {
		return Product{}, errors.New("can't find product category")
	}

	product := Product{
		AggreateRoot: AggreateRoot{
			ID: guid.NewString(),
		},
		Name:          name,
		Description:   descripion,
		Active:        true,
		UpdateAt:      DateNow(),
		CreateAt:      DateNow(),
		Skus:          []Sku{},
		CategoriesIds: categoriesIDs,
	}

	return product, nil
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

//FindSku search sku
func (p *Product) FindSku(skuID string) (Sku, error) {
	for _, item := range p.Skus {
		if skuID == item.ID {
			return item, nil
		}
	}

	return Sku{}, errors.New("sku not found")
}

//UpdateSku update sku
func (p *Product) UpdateSku(skuToUpdate Sku) error {
	found := false

	for i, sku := range p.Skus {
		if sku.ID == skuToUpdate.ID {
			p.Skus[i] = skuToUpdate
			found = true
		}
	}

	if found == false {
		return errors.New("Sku not found")
	}

	p.UpdateAt = DateNow()

	return nil
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
func (p *Product) AddCategory(categoryID string) {
	p.CategoriesIds = append(p.CategoriesIds, categoryID)
	p.UpdateAt = DateNow()
}

//RemoveCategory remove product category
func (p *Product) RemoveCategory(categoryID string) error {
	index := sort.Search(len(p.CategoriesIds), func(i int) bool {
		return p.CategoriesIds[i] == categoryID
	})

	if len(p.CategoriesIds) == index {
		return errors.New("category not found")
	}

	p.CategoriesIds = append(p.CategoriesIds[:index], p.CategoriesIds[index+1:]...)
	p.UpdateAt = DateNow()
	return nil
}
