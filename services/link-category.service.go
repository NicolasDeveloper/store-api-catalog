package services

import (
	"errors"

	"github.com/NicolasDeveloper/store-catalog-api/domain"
)

//LinkCategory link categories
func LinkCategory(
	categoryID string,
	category domain.Category,
	categoryRepository domain.CategoryRepository) (domain.Category, error) {

	if categoryID == "" {
		return category, nil
	}

	parentCategory, err := categoryRepository.GetByID(categoryID)

	if err != nil {
		return category, errors.New("category to link not found")
	}

	category.Link(parentCategory.ID)

	return category, nil
}
