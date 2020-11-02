package infra

import (
	"context"

	"github.com/NicolasDeveloper/store-catalog-api/domain"
	"go.mongodb.org/mongo-driver/bson"
)

type productRepository struct {
	dbctx *DbContext
	domain.ProductRepository
}

//NewProductRepository contructor
func NewProductRepository(dbctx *DbContext) domain.ProductRepository {
	repo := productRepository{
		dbctx: dbctx,
	}

	return &repo
}

func (r *productRepository) Save(product domain.Product) error {
	collection, error := r.dbctx.GetCollection(product)
	_, error = collection.InsertOne(context.TODO(), product)
	return error
}

func (r *productRepository) Update(product domain.Product) error {
	collection, error := r.dbctx.GetCollection(product)

	filter := bson.M{"_id": product.ID}

	_, error = collection.UpdateOne(context.TODO(), filter, bson.M{"$set": product})
	return error
}

func (r *productRepository) GetByID(productID string) (domain.Product, error) {
	collection, error := r.dbctx.GetCollection(domain.Product{})
	filter := bson.M{"_id": productID}

	var product domain.Product

	error = collection.FindOne(context.TODO(), filter).Decode(&product)

	return product, error
}
