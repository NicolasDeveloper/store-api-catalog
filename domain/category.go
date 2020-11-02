package domain

//Category product cateogory
type Category struct {
	Name             string `bson:"name" json:"name"`
	ParentCategoryID string `bson:"parent_category_id" json:"parent_category_id"`
	Entity
}
