package domain

import (
	"github.com/beevik/guid"
)

//ParentCategory parent
type ParentCategory *Category

//Category product cateogory
type Category struct {
	Name             string `bson:"name" json:"name"`
	Path             string `bson:"path" json:"path"`
	ParentCategoryID string `bson:"parent_category_id" json:"parent_category_id"`
	AggreateRoot
}

//NewCategory constructor
func NewCategory(name string, path string) (Category, error) {
	return Category{
		AggreateRoot: AggreateRoot{
			ID: guid.NewString(),
		},
		Name: name,
		Path: path,
	}, nil
}

//SetName update name
func (c *Category) SetName(name string) {
	c.Name = name
}

//SetPath update name
func (c *Category) SetPath(path string) {
	c.Path = path
}

//Link link parent with child category
func (c *Category) Link(categoryID string) {
	c.ParentCategoryID = categoryID
}
