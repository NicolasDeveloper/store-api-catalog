package domain

//Category product cateogory
type Category struct {
	ID               string `bson:"_id" json:"id"`
	Name             string `bson:"name" json:"name"`
	ParentCategoryID string `bson:"parent_category_id" json:"parent_category_id"`
	IEntity
}

//GetID get identification
func (c *Category) GetID() string {
	return c.ID
}
