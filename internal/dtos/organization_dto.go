package dtos

import "github.com/asyncnavi/raateo/internal/models"

type CreateOrganizationRequest struct {
	Name        string `json:"name" validate:"required,min=2,max=256"`
	Description string `json:"description" validate:"max=2000"`
	LogoURL     string `json:"logo_url"`
}

type OrganizationResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      uint   `json:"user_id"`
	LogoURL     string `json:"logo_url"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func GenerateOrganizationResponse(organization *models.Organization) *OrganizationResponse {
	return &OrganizationResponse{
		ID:          organization.ID,
		Name:        organization.Name,
		UserID:      organization.UserID,
		Description: organization.Description,
		LogoURL:     organization.LogoURL,
		CreatedAt:   organization.CreatedAt.String(),
		UpdatedAt:   organization.UpdatedAt.String(),
	}
}
