package domain

//ProductRepository repo
type ProductRepository interface {
	Save(product Product) error
}
