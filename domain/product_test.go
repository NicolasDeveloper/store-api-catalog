package domain

import (
	"testing"

	"github.com/beevik/guid"

	"github.com/stretchr/testify/assert"
)

func TestProduct(t *testing.T) {
	category := []Category{Category{ID: guid.NewString(), Name: "Teste", ParentCategoryID: ""}}

	t.Run("Should create priduct", func(t *testing.T) {
		assert := assert.New(t)
		product, err := NewProduct("teste", "description test", category)

		assert.Empty(err)
		assert.NotEmpty(product.ID)
	})

	t.Run("Should create a guid identification", func(t *testing.T) {
		assert := assert.New(t)
		product, err := NewProduct("teste", "description test", category)

		assert.Empty(err)
		assert.True(guid.IsGuid(product.ID))
	})

	t.Run("Should input date when create a product", func(t *testing.T) {
		assert := assert.New(t)
		product, err := NewProduct("teste", "description test", category)

		assert.Empty(err)
		assert.NotEmpty(product.CreateAt)
	})

	t.Run("Should create active product", func(t *testing.T) {
		assert := assert.New(t)
		product, err := NewProduct("teste", "description test", category)

		assert.Empty(err)
		assert.True(product.Active)
	})

	t.Run("Should inactive product", func(t *testing.T) {
		assert := assert.New(t)
		product, err := NewProduct("teste", "description test", category)
		product.Disable()

		assert.Empty(err)
		assert.False(product.Active)
	})

	t.Run("Should remove category", func(t *testing.T) {
		assert := assert.New(t)
		product, err := NewProduct("teste", "description test", category)
		assert.Empty(err)

		c := Category{
			ID:   "xpto",
			Name: "Teste remove",
		}

		product.AddCategory(c)
		assert.Equal(len(product.Categories), 2)

		err = product.RemoveCategory("xpto")
		assert.Equal(len(product.Categories), 1)
		assert.Empty(err)
	})
}
