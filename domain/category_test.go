package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCsategory(t *testing.T) {
	t.Run("Should create category", func(t *testing.T) {
		assert := assert.New(t)

		c, _ := NewCategory("Teste de categoria", "teste")

		assert.NotEmpty(t, c.ID)
	})

	t.Run("Should link category", func(t *testing.T) {
		assert := assert.New(t)

		c, _ := NewCategory("Teste de categoria", "teste")
		b, _ := NewCategory("Teste de categoria", "teste")

		b.Link(c.ID)

		assert.Equal(b.ParentCategoryID, c.ID)
	})
}
