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

	"github.com/gin-gonic/gin"
)

// GetSellerProductDetail handles GET /v1/seller/products/:id
// Returns full product detail for seller including galleries
func GetSellerProductDetail(appCtx appctx.AppContext) gin.HandlerFunc {
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

		// Setup dependencies
		productStorage := productstorage.NewSQLStore(db)
		sellerStorage := sellerstorage.NewSQLStore(db)

		sfinder := sellerrepository.NewGetSellerRepo(sellerStorage)
		pfinder := productrepository.NewFindProductRepo(productStorage)
		galleriesRepo := productrepository.NewGetGalleriesRepo(productStorage)
		categoryRepo := productrepository.NewGetProductCategoryRepo(productStorage)

		biz := productbusiness.NewGetSellerProductDetailBusiness(
			sfinder,
			pfinder,
			galleriesRepo,
			categoryRepo,
		)

		result, err := biz.GetSellerProductDetail(c.Request.Context(), userID, productID)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
