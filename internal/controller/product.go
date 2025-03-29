package controller

import (
	"errors"
	"github.com/asyncnavi/raateo/internal/dtos"
	apiErrors "github.com/asyncnavi/raateo/pkg/errros"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log/slog"
	"net/http"
)

type ProductParams struct {
	ProductID      uint `uri:"product_id" binding:"required"`
	OrganizationID uint `uri:"organization_id" binding:"required"`
}

func (ctrl *Controller) GetProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params ProductParams

		if err := c.ShouldBindUri(&params); err != nil {
			apiErrors.RespondWithError(c, err)
			return
		}

		product, err := ctrl.repo.GetProduct(params.ProductID, params.OrganizationID)

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				slog.Error("Product not found", slog.Any("product_id", params.ProductID), slog.Any("organization_id", params.OrganizationID))
				apiErrors.RespondWithError(c, errors.New("product not found"))
				return
			}
			slog.Error("Failed to get product", slog.Any("product_id", params.ProductID), slog.Any("organization_id", params.OrganizationID), slog.Any("error", err))
			apiErrors.RespondWithError(c, err)
			return
		}

		result := dtos.GenerateProductResponse(&product)
		c.JSON(http.StatusOK, result)
	}
}

func (ctrl *Controller) GetProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		products, err := ctrl.repo.GetProducts()
		if err != nil {
			slog.Error("failed to fetch products")
			apiErrors.RespondWithError(c, err)
			return
		}
		result := dtos.GenerateProductsResponse(products)
		c.JSON(http.StatusOK, result)
	}
}

func (ctrl *Controller) GetProductsForOrganization() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		user := UserFromContext(ctx)
		org, err := ctrl.repo.OrganizationFromUser(user.ID)

		if err != nil {
			apiErrors.RespondWithError(c, errors.New("failed to parse organization id"))
			return
		}
		products, err := ctrl.repo.GetProductsFromOrganizationId(org.ID)

		if err != nil && errors.Is(gorm.ErrRecordNotFound, err) {
			apiErrors.RespondWithError(c, errors.New("no products found for the organization"))
			return
		}

		if err != nil {
			apiErrors.InternalError()
			return
		}

		result := dtos.GenerateProductsResponse(products)
		c.JSON(http.StatusOK, result)
	}
}

func (ctrl *Controller) CreateProductForOrganization() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		var reqInput dtos.CreateProductRequest

		if err := c.ShouldBindJSON(&reqInput); err != nil {
			apiErrors.RespondWithError(c, err)
			return
		}

		user := UserFromContext(ctx)
		org, err := ctrl.repo.OrganizationFromUser(user.ID)

		if err != nil {
			apiErrors.RespondWithError(c, errors.New("failed to parse organization id"))
			return
		}

		product := dtos.GenerateCreateProductDTO(&reqInput, org.ID)

		if err := ctrl.repo.SaveProduct(product); err != nil {
			apiErrors.InternalError()
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Product is saved",
		})
	}
}

//func (ctrl *Controller) CreateProduct() gin.HandlerFunc {
//	var reqArgs struct {
//		OrganizationID string `json:"organization_id" validate:"required"`
//		Name           string `json:"name" validate:"required, min=2, max=256"`
//		Description    string `json:"description" validate:"max=2000"`
//		LogoURL        string `json:"logo_url"`
//		ThumbnailURL   string `json:"thumbnail_url" `
//	}
//	return func(c *gin.Context) {
//
//		if err := c.ShouldBindJSON(&reqArgs); err != nil {
//			slog.Error("failed to parse json", "", err)
//			apiErrors.RespondWithError(c, err)
//			return
//		}
//
//		// Parse Organization ID to uint
//		orgID, err := strconv.Atoi(reqArgs.OrganizationID)
//
//		if err != nil {
//			slog.Error("failed to parse organization_id", "", err)
//			apiErrors.RespondWithError(c, err)
//			return
//		}
//
//		prod := &database.Product{
//			OrganizationID: uint(orgID),
//			Name:           reqArgs.Name,
//			Description:    reqArgs.Description,
//			LogoURL:        reqArgs.LogoURL,
//			ThumbnailURL:   reqArgs.ThumbnailURL,
//		}
//		if err := ctrl.db.SaveProduct(prod); err != nil {
//			slog.Error("failed to save product")
//			apiErrors.RespondWithError(c, err)
//			return
//		}
//
//		resultProduct := gin.H{
//			"organization_id": uint(orgID),
//			"name":            reqArgs.Name,
//			"description":     reqArgs.Description,
//			"logo_url":        reqArgs.LogoURL,
//			"thumbnail_url":   reqArgs.ThumbnailURL,
//		}
//
//		c.JSON(http.StatusOK, resultProduct)
//	}
//}
//
//func (ctrl *Controller) ListProducts() gin.HandlerFunc {
//
//	return func(c *gin.Context) {
//		products, err := ctrl.db.GetAllProducts()
//		if err != nil {
//			slog.Error("failed to fetch products")
//			apiErrors.RespondWithError(c, err)
//			return
//		}
//
//		var resultProducts []map[string]interface{}
//
//		for _, product := range products {
//			resultProducts = append(resultProducts, map[string]interface{}{
//				"id":              product.ID,
//				"organization_id": product.OrganizationID,
//				"name":            product.Name,
//				"description":     product.Description,
//				"logo_url":        product.LogoURL,
//				"thumbnail_url":   product.ThumbnailURL,
//				"created_at":      product.CreatedAt,
//				"updated_at":      product.UpdatedAt,
//				"deleted_at":      product.DeletedAt,
//			})
//		}
//
//		c.JSON(http.StatusOK, resultProducts)
//	}
//}
//
//func (ctrl *Controller) ListOrganizationProducts() gin.HandlerFunc {
//
//	return func(c *gin.Context) {
//		// get the id
//		id := c.Param("org_id")
//
//		// parsing product id
//		orgId, err := strconv.Atoi(id)
//
//		if err != nil {
//			slog.Error("failed to parse organization Id")
//			apiErrors.RespondWithError(c, err)
//			return
//		}
//
//		products, err := ctrl.db.GetProductsByOrganization(uint(orgId))
//
//		if err != nil {
//			slog.Error("failed to fetch products")
//			apiErrors.RespondWithError(c, err)
//			return
//		}
//
//		var resultProducts []map[string]interface{}
//
//		for _, product := range products {
//			resultProducts = append(resultProducts, map[string]interface{}{
//				"id":              product.ID,
//				"organization_id": product.OrganizationID,
//				"name":            product.Name,
//				"description":     product.Description,
//				"logo_url":        product.LogoURL,
//				"thumbnail_url":   product.ThumbnailURL,
//				"created_at":      product.CreatedAt,
//				"updated_at":      product.UpdatedAt,
//				"deleted_at":      product.DeletedAt,
//			})
//		}
//
//		c.JSON(http.StatusOK, resultProducts)
//	}
//}
