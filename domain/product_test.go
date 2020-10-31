package domain

import (
	"testing"

	"github.com/beevik/guid"

	"github.com/stretchr/testify/assert"
)

func TestProduct(t *testing.T) {
	t.Run("Should create priduct", func(t *testing.T) {
		assert := assert.New(t)
		product, err := NewProduct("teste", "description test")

		assert.Empty(err)
		assert.NotEmpty(product.ID)
	})

	t.Run("Should create a guid identification", func(t *testing.T) {
		assert := assert.New(t)
		product, err := NewProduct("teste", "description test")

		assert.Empty(err)
		assert.True(guid.IsGuid(product.ID))
	})

	t.Run("Should input date when create a product", func(t *testing.T) {
		assert := assert.New(t)
		product, err := NewProduct("teste", "description test")

		assert.Empty(err)
		assert.NotEmpty(product.CreateAt)
	})

	t.Run("Should create active product", func(t *testing.T) {
		assert := assert.New(t)
		product, err := NewProduct("teste", "description test")

		assert.Empty(err)
		assert.True(product.Active)
	})

	t.Run("Should inactive product", func(t *testing.T) {
		assert := assert.New(t)
		product, err := NewProduct("teste", "description test")
		product.Disable()

		assert.Empty(err)
		assert.False(product.Active)
	})

	t.Run("Should update product", func(t *testing.T) {
		assert := assert.New(t)
		product, err := NewProduct("teste", "description test")

		product.Update("updating product name", "updating product description")

		assert.Empty(err)
		assert.Equal(product.Name, "updating product name")
		assert.Equal(product.Description, "updating product description")
	})

	t.Run("Should update time when product is updated with new information", func(t *testing.T) {
		assert := assert.New(t)
		product, err := NewProduct("teste", "description test")

		now := DateNow()
		product.Update("updating product name", "updating product description")

		assert.Empty(err)
		expetedDiff := product.UpdateAt.Sub(now).Seconds()

		assert.True(expetedDiff < 1)
	})
}
