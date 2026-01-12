package productgin

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	productbusiness "OpenMarket/module/product/business"
	productrepository "OpenMarket/module/product/repository"
	productstorage "OpenMarket/module/product/storage"
	sellerrepository "OpenMarket/module/seller/repository"
	sellerstorage "OpenMarket/module/seller/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ToggleStatusRequest struct {
	Status *int `json:"status" binding:"required"`
}

// ToggleVariantStatus handles PATCH /v1/seller/products/:id/variant/:vid/status
func ToggleVariantStatus(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		u, ok := c.MustGet(common.CurrentUser).(common.Requester)
		if !ok {
			panic(common.ErrUnauthorized(nil))
		}
		userID := u.GetUserId()

		// Parse product ID
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.InvalidRequestError(err))
		}
		productID := int(uid.GetLoacalID())

		// Parse variant ID
		variantID, err := strconv.Atoi(c.Param("vid"))
		if err != nil {
			panic(common.InvalidRequestError(err))
		}

		// Parse request body
		var req ToggleStatusRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			panic(common.InvalidRequestError(err))
		}

		storage := productstorage.NewSQLStore(db)
		sellerStorage := sellerstorage.NewSQLStore(db)
		sfinder := sellerrepository.NewGetSellerRepo(sellerStorage)
		pfinder := productrepository.NewFindProductRepo(storage)
		statusRepo := productrepository.NewUpdateVariantStatusRepo(storage)
		variantReader := productrepository.NewVariantReaderRepo(storage)

		biz := productbusiness.NewToggleVariantStatusBusiness(sfinder, pfinder, statusRepo, variantReader)

		if err := biz.ToggleStatus(c.Request.Context(), userID, productID, variantID, *req.Status); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
