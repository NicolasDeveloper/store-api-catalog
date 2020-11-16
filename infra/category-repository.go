package infra

import (
	"context"

	"github.com/NicolasDeveloper/store-catalog-api/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type categoryRepository struct {
	dbctx *DbContext
	domain.CategoryRepository
}

//NewCategoryRepository contructor
func NewCategoryRepository(dbctx *DbContext) domain.CategoryRepository {
	return &categoryRepository{
		dbctx: dbctx,
	}
}

func (r *categoryRepository) Save(category domain.Category) error {
	collection, error := r.dbctx.GetCollection(category)
	_, error = collection.InsertOne(context.TODO(), category)
	return error
}

func (r *categoryRepository) Update(category domain.Category) error {
	collection, error := r.dbctx.GetCollection(category)

	filter := bson.M{"_id": category.ID}

	_, error = collection.UpdateOne(context.TODO(), filter, bson.M{"$set": category})
	return error
}

func (r *categoryRepository) GetByID(categoryID string) (domain.Category, error) {
	collection, error := r.dbctx.GetCollection(domain.Category{})
	filter := bson.M{"_id": categoryID}

	var category domain.Category

	error = collection.FindOne(context.TODO(), filter).Decode(&category)

	return category, error
}

func (r *categoryRepository) GetCategories(categoryIDs []string) ([]domain.Category, error) {
	collection, error := r.dbctx.GetCollection(domain.Category{})
	filter := bson.M{"_id": bson.M{"$in": categoryIDs}}

	var categories []domain.Category

	cursor, error := collection.Find(context.TODO(), filter)

	if error != nil {
		return nil, error
	}

	cursor.All(context.Background(), &categories)

	return categories, error
}

func (r *categoryRepository) ListCategories(skip int, take int) ([]domain.Category, error) {
	collection, error := r.dbctx.GetCollection(domain.Category{})

	var categories []domain.Category

	options := options.FindOptions{}

	options.SetSkip(int64(skip))
	options.SetLimit(int64(take))

	cursor, error := collection.Find(context.TODO(), bson.D{}, &options)

	if error != nil {
		return nil, error
	}

	cursor.All(context.Background(), &categories)

	return categories, error
}

func (r *categoryRepository) Total() (int, error) {
	collection, error := r.dbctx.GetCollection(domain.Category{})
	itemsCount, error := collection.CountDocuments(context.TODO(), bson.D{})

	return int(itemsCount), error
}
