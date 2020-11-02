package domain

//CategoryRepository repo
type CategoryRepository interface {
	Save(category Category) error
	Update(category Category) error
	GetByID(categoryID string) (Category, error)
	GetCategories(categoriesIDs []string) ([]Category, error)
}
