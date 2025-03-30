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

func (ctrl *Controller) DeleteProductForOrganization() gin.HandlerFunc {

	var params struct {
		ProductID uint `uri:"product_id" binding:"required"`
	}
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		if err := c.ShouldBindUri(&params); err != nil {
			apiErrors.RespondWithError(c, err)
			return
		}

		user := UserFromContext(ctx)
		org, err := ctrl.repo.OrganizationFromUser(user.ID)

		if err != nil {
			apiErrors.RespondWithError(c, errors.New("failed to parse organization id"))
			return
		}

		if err := ctrl.repo.DeleteProduct(params.ProductID, org.ID); err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				apiErrors.RespondWithError(c, errors.New("product not found"))
				return
			}
			apiErrors.InternalError()
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Product deleted successfully",
		})
	}
}

func (ctrl *Controller) SearchProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		query := c.Query("q")
		if query == "" {
			apiErrors.RespondWithError(c, errors.New("query parameter is required"))
			return
		}

		products, err := ctrl.repo.GetProductWithQuery(query)
		if err != nil {
			slog.Error("failed to search products", slog.Any("query", query), slog.Any("error", err))
			apiErrors.RespondWithError(c, err)
			return
		}

		result := dtos.GenerateProductsResponse(products)
		c.JSON(http.StatusOK, result)
	}
}
