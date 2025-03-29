package dtos

import (
	"github.com/asyncnavi/raateo/internal/models"
)

type RatingResponse struct {
	ID             uint   `json:"id"`
	OrganizationID uint   `json:"organization_id" validate:"required"`
	RatingValue    uint   `json:"rating_value" validate:"required"`
	UserID         uint   `json:"user_id" validate:"required"`
	Username       string `json:"username"`
	Comment        string `json:"comment"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

type CreateRatingRequest struct {
	OrganizationID uint   `json:"organization_id" validate:"required"`
	UserID         uint   `json:"user_id" validate:"required"`
	ProductID      uint   `json:"product_id" validate:"required"`
	FeatureID      uint   `json:"feature_id" validate:"required"`
	RatingValue    uint   `json:"rating_value" validate:"required"`
	Comment        string `json:"comment,omitempty"`
}

func GenerateCreateRatingDTO(rating *CreateRatingRequest) *models.Rating {
	return &models.Rating{
		OrganizationID: rating.OrganizationID,
		ProductID:      rating.ProductID,
		FeatureID:      rating.FeatureID,
		RatingValue:    rating.RatingValue,
		UserID:         rating.UserID,
		Comment:        rating.Comment,
	}
}

func GenerateRatingResponse(rating *models.Rating) RatingResponse {
	return RatingResponse{
		ID:             rating.ID,
		OrganizationID: rating.OrganizationID,
		RatingValue:    rating.RatingValue,
		UserID:         rating.UserID,
		Comment:        rating.Comment,
		CreatedAt:      rating.CreatedAt.String(),
		UpdatedAt:      rating.UpdatedAt.String(),
	}
}

func GenerateRatingsResponse(ratings []models.Rating) []RatingResponse {

	var response []RatingResponse
	for _, rating := range ratings {
		response = append(response, RatingResponse{
			ID:             rating.ID,
			OrganizationID: rating.OrganizationID,
			RatingValue:    rating.RatingValue,
			CreatedAt:      rating.CreatedAt.String(),
			Username:       rating.User.FirstName,
			Comment:        rating.Comment,
			UpdatedAt:      rating.UpdatedAt.String(),
		})
	}
	return response
}
