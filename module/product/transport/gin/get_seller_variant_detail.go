package productgin

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	productbusiness "OpenMarket/module/product/business"
	productmodel "OpenMarket/module/product/model"
	productrepository "OpenMarket/module/product/repository"
	productstorage "OpenMarket/module/product/storage"
	sellerrepository "OpenMarket/module/seller/repository"
	sellerstorage "OpenMarket/module/seller/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetSellerVariantDetail handles GET /v1/seller/products/:id/variant/:vid
// Returns full variant detail for seller including attribute IDs
func GetSellerVariantDetail(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		// Get current user
		u, ok := c.MustGet(common.CurrentUser).(common.Requester)
		if !ok {
			panic(common.ErrUnauthorized(nil))
		}
		userID := u.GetUserId()

		// Parse product ID
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrEntityNotFound(productmodel.EntityName, err))
		}
		productID := int(uid.GetLoacalID())

		// Parse variant ID
		variantID, err := strconv.Atoi(c.Param("vid"))
		if err != nil {
			panic(common.ErrEntityNotFound("Variant", err))
		}

		// Setup dependencies
		productStorage := productstorage.NewSQLStore(db)
		sellerStorage := sellerstorage.NewSQLStore(db)

		sfinder := sellerrepository.NewGetSellerRepo(sellerStorage)
		pfinder := productrepository.NewFindProductRepo(productStorage)
		variantRepo := productrepository.NewGetVariantWithAttrsRepo(productStorage)

		biz := productbusiness.NewGetSellerVariantDetailBusiness(
			sfinder,
			pfinder,
			variantRepo,
		)

		result, err := biz.GetSellerVariantDetail(c.Request.Context(), userID, productID, variantID)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
