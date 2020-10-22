package domain

//ProductRepository repo
type ProductRepository interface {
	Save(product Product) error
	Update(product Product) error
	GetByID(productID string) (Product, error)
}
