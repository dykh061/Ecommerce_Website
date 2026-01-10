package cartgin

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	cartbusiness "OpenMarket/module/cart/business"
	cartstorage "OpenMarket/module/cart/storage"

	"github.com/gin-gonic/gin"
)

type RemoveItemRequest struct {
	VariantId int `json:"variant_id" binding:"required"`
}

func RemoveItemFromCart(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := appCtx.GetMainDBConnection()
		u, ok := context.MustGet(common.CurrentUser).(common.Requester)
		if !ok {
			panic(common.ErrUnauthorized(nil))
		}
		userId := u.GetUserId()

		var req RemoveItemRequest
		if err := context.ShouldBindJSON(&req); err != nil {
			panic(common.InvalidRequestError(err))
		}

		if req.VariantId <= 0 {
			panic(common.InvalidRequestError(nil))
		}

		storage := cartstorage.NewSQLStore(db)
		biz := cartbusiness.NewRemoveItemFromCartBusiness(storage)

		if err := biz.RemoveItemFromCart(context.Request.Context(), userId, req.VariantId); err != nil {
			panic(err)
		}

		context.JSON(200, common.SimpleSuccessResponse(true))
	}
}
