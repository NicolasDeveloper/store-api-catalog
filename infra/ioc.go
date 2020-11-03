package infra

import (
	"github.com/NicolasDeveloper/store-catalog-api/domain"
	"github.com/golobby/container"
)

type ioc struct {
}

//NewContainer initialize container
func NewContainer() {
	c := &ioc{}

	c.registerDependencies()
}

func (i *ioc) registerDependencies() {
	container.Singleton(func() *DbContext {
		return NewDbContext()
	})

	container.Transient(func(dbctx *DbContext) domain.ProductRepository {
		return NewProductRepository(dbctx)
	})

	container.Transient(func(dbctx *DbContext) domain.CategoryRepository {
		return NewCategoryRepository(dbctx)
	})
}
