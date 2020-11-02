package domain

import (
	"testing"

	"github.com/beevik/guid"

	"github.com/stretchr/testify/assert"
)

func TestProduct(t *testing.T) {
	category := []Category{
		Category{
			AggreateRoot: AggreateRoot{
				ID: guid.NewString(),
			},
			Name: "Teste",
		},
	}

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
			AggreateRoot: AggreateRoot{
				ID: "xpto",
			},
			Name: "Teste remove",
		}

		product.AddCategory(c)
		assert.Equal(len(product.Categories), 2)

		err = product.RemoveCategory("xpto")
		assert.Equal(len(product.Categories), 1)
		assert.Empty(err)
	})

	t.Run("Should add sku", func(t *testing.T) {
		assert := assert.New(t)
		product, err := NewProduct("teste", "description test", category)
		assert.Empty(err)

		s := NewSku("Teste sku", 20.00, 20)

		product.AddSku(s)
		assert.Equal(len(product.Skus), 1)

		assert.Empty(err)
	})

	t.Run("Should remove sku", func(t *testing.T) {
		assert := assert.New(t)
		product, err := NewProduct("teste", "description test", category)
		assert.Empty(err)

		s := NewSku("Teste sku", 20.00, 20)

		product.AddSku(s)
		assert.Equal(len(product.Skus), 1)

		product.RemoveSku(s.ID)
		assert.Equal(len(product.Skus), 0)

		assert.Empty(err)
	})
}
