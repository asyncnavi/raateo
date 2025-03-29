package dtos

import (
	"github.com/asyncnavi/raateo/internal/models"
)

type CreateProductRequest struct {
	Name         string `json:"name" validate:"required,min=2,max=256"`
	Description  string `json:"description" validate:"max=2000"`
	LogoURL      string `json:"logo_url"`
	ThumbnailURL string `json:"thumbnail_url"`
	WebsiteURL   string `json:"website_url"`
}

type ProductResponse struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description,omitempty"`
	LogoURL        string `json:"logo_url,omitempty"`
	ThumbnailURL   string `json:"thumbnail_url,omitempty"`
	OrganizationID uint   `json:"organization_id"`
	UserID         uint   `json:"user_id"`
	WebsiteURL     string `json:"website_url,omitempty"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

func GenerateProductsResponse(products []models.Product) []ProductResponse {
	var productsResponse []ProductResponse
	for _, product := range products {
		productsResponse = append(productsResponse, ProductResponse{
			ID:             product.ID,
			OrganizationID: product.OrganizationID,
			Name:           product.Name,
			Description:    product.Description,
			LogoURL:        product.LogoURL,
			CreatedAt:      product.CreatedAt.String(),
			UpdatedAt:      product.UpdatedAt.String(),
		})
	}
	return productsResponse
}

func GenerateProductResponse(product *models.Product) ProductResponse {
	return ProductResponse{
		ID:             product.ID,
		Name:           product.Name,
		Description:    product.Description,
		OrganizationID: product.OrganizationID,
		LogoURL:        product.LogoURL,
		CreatedAt:      product.CreatedAt.String(),
		UpdatedAt:      product.UpdatedAt.String(),
	}
}

func GenerateCreateProductDTO(product *CreateProductRequest, organizationID uint) *models.Product {
	return &models.Product{
		Name:           product.Name,
		Description:    product.Description,
		LogoURL:        product.LogoURL,
		ThumbnailURL:   product.ThumbnailURL,
		WebsiteURL:     product.WebsiteURL,
		OrganizationID: organizationID,
	}
}
