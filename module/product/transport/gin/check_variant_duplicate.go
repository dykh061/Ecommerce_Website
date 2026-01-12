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

	"github.com/gin-gonic/gin"
)

type CheckDuplicateRequest struct {
	AttributeValueIDs []int `json:"attribute_value_ids" binding:"required"`
	ExcludeVariantID  *int  `json:"exclude_variant_id"`
}

type CheckDuplicateResponse struct {
	Exists bool `json:"exists"`
}

// CheckVariantDuplicate handles POST /v1/seller/products/:id/variants/check-duplicate
func CheckVariantDuplicate(appCtx appctx.AppContext) gin.HandlerFunc {
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

		// Parse request body
		var req CheckDuplicateRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			panic(common.InvalidRequestError(err))
		}

		storage := productstorage.NewSQLStore(db)
		sellerStorage := sellerstorage.NewSQLStore(db)
		sfinder := sellerrepository.NewGetSellerRepo(sellerStorage)
		pfinder := productrepository.NewFindProductRepo(storage)
		duplicateRepo := productrepository.NewVariantDuplicateRepo(storage)

		biz := productbusiness.NewCheckVariantDuplicateBusiness(sfinder, pfinder, duplicateRepo)

		exists, err := biz.CheckDuplicate(c.Request.Context(), userID, productID, req.AttributeValueIDs, req.ExcludeVariantID)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(CheckDuplicateResponse{Exists: exists}))
	}
}
