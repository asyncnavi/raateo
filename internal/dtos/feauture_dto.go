package dtos

import "github.com/asyncnavi/raateo/internal/models"

type CreateFeatureRequest struct {
	OrganizationID uint   `json:"organization_id" validate:"required"`
	ProductID      uint   `json:"product_id" validate:"required"`
	Name           string `json:"name" validate:"required,min=2,max=256"`
	Description    string `json:"description" validate:"max=2000"`
	VideoURL       string `json:"video_url"`
	ThumbnailURL   string `json:"thumbnail_url"`
}

type FeatureResponse struct {
	ID             uint   `json:"id"`
	OrganizationID uint   `json:"organization_id"`
	ProductID      uint   `json:"product_id"`
	Name           string `json:"name"`
	Description    string `json:"description,omitempty"`
	VideoURL       string `json:"video_url,omitempty"`
	ThumbnailURL   string `json:"thumbnail_url,omitempty"`
	Status         string `json:"status,omitempty"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

func GenerateFeaturesResponse(features []models.Feature) []FeatureResponse {
	var featuresResponse []FeatureResponse
	for _, feature := range features {
		featuresResponse = append(featuresResponse, FeatureResponse{
			ID:             feature.ID,
			OrganizationID: feature.OrganizationID,
			ProductID:      feature.ProductID,
			Name:           feature.Name,
			Description:    feature.Description,
			VideoURL:       feature.VideoUrl,
			ThumbnailURL:   feature.ThumbnailUrl,
			Status:         string(feature.Status),
			CreatedAt:      feature.CreatedAt.String(),
			UpdatedAt:      feature.UpdatedAt.String(),
		})
	}
	return featuresResponse
}

func GenerateFeatureResponse(feature *models.Feature) FeatureResponse {
	return FeatureResponse{
		ID:             feature.ID,
		OrganizationID: feature.OrganizationID,
		ProductID:      feature.ProductID,
		Name:           feature.Name,
		Description:    feature.Description,
		VideoURL:       feature.VideoUrl,
		ThumbnailURL:   feature.ThumbnailUrl,
		Status:         string(feature.Status),
		CreatedAt:      feature.CreatedAt.String(),
		UpdatedAt:      feature.UpdatedAt.String(),
	}
}
