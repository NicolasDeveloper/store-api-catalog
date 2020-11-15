package domain

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type skipTestCase struct {
	PageIndex    int
	ExpectedPage int
}

func TestCategories(t *testing.T) {

	t.Run("Should create a paginator", func(t *testing.T) {
		assert := assert.New(t)

		_, err := NewPaginator(0, 10, 30)

		assert.Empty(err)
	})

	t.Run("Should next page index be equals total items when page current page index is bigger than total items", func(t *testing.T) {
		assert := assert.New(t)

		paginator, err := NewPaginator(350, 10, 300)

		assert.Equal(300, paginator.GetNextPageIndex())
		assert.Empty(err)
	})

	t.Run("Should next page index be equals total items when page size is bigger than total items", func(t *testing.T) {
		assert := assert.New(t)

		paginator, err := NewPaginator(300, 10, 300)

		assert.Equal(300, paginator.GetNextPageIndex())
		assert.Empty(err)
	})

	t.Run("Should get next skip items", func(t *testing.T) {
		assert := assert.New(t)

		paginator, err := NewPaginator(200, 10, 300)

		assert.Equal(210, paginator.GetNextPageIndex())
		assert.Empty(err)
	})

	skipCases := []skipTestCase{
		skipTestCase{
			PageIndex:    0,
			ExpectedPage: 1,
		},
		skipTestCase{
			PageIndex:    200,
			ExpectedPage: 2,
		},
		skipTestCase{
			PageIndex:    451,
			ExpectedPage: 4,
		},
	}

	for _, skipCase := range skipCases {
		t.Run("Should get next skip items from page index "+strconv.Itoa(skipCase.PageIndex), func(t *testing.T) {
			assert := assert.New(t)

			paginator, err := NewPaginator(skipCase.PageIndex, 100, 400)

			assert.Equal(skipCase.ExpectedPage, paginator.GetNextPage())
			assert.Empty(err)
		})
	}
}
