package productgin

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	productbusiness "OpenMarket/module/product/business"
	productmodel "OpenMarket/module/product/model"
	productrepository "OpenMarket/module/product/repository"
	productstorage "OpenMarket/module/product/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetProductAttributes handles GET /v1/products/:id/attributes
// Returns list of attributes and their values for a product (public API)
func GetProductAttributes(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		// Parse product ID
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrEntityNotFound(productmodel.EntityName, err))
		}
		productID := int(uid.GetLoacalID())

		// Setup dependencies
		storage := productstorage.NewSQLStore(db)

		productFinder := productrepository.NewGetProductRepo(storage)
		attrRepo := productrepository.NewGetProductAttributesRepo(storage)

		biz := productbusiness.NewGetProductAttributesBusiness(
			productFinder,
			attrRepo,
		)

		result, err := biz.GetProductAttributes(c.Request.Context(), productID)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
