package infra

import (
	"context"

	"github.com/NicolasDeveloper/store-catalog-api/domain"
)

type repository struct {
	dbctx *DbContext
	domain.ProductRepository
}

//NewProductRepository contructor
func NewProductRepository(dbctx *DbContext) domain.ProductRepository {
	return &repository{
		dbctx: dbctx,
	}
}

//Save save product
func (r *repository) Save(product domain.Product) error {
	collection, error := r.dbctx.GetCollection(product)
	_, error = collection.InsertOne(context.TODO(), product)
	return error
}
