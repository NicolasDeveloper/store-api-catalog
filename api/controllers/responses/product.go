package responses

import (
	"github.com/NicolasDeveloper/store-catalog-api/domain"
)

// ProductResponse product view model
type ProductResponse struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Skus        []SkuResponse `json:"skus"`
}

// AddSkus from sku to sku response
func (p *ProductResponse) AddSkus(skus []domain.Sku) {
	p.Skus = make([]SkuResponse, len(skus))

	for _, sku := range skus {
		skuResponse := SkuResponse{
			ID:       sku.ID,
			Name:     sku.Name,
			Price:    sku.Price,
			Quantity: sku.Quantity,
		}

		p.Skus = append(p.Skus, skuResponse)
	}
}

// ResponseData response payload
//swagger:response productResponse
type swaggerproductResponse struct {
	//in: body
	Body struct {
		Success bool            `json:"success"`
		Data    ProductResponse `json:"data"`
	}
}
