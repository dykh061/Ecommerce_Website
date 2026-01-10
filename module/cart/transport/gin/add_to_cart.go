package cartgin

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	cartbusiness "OpenMarket/module/cart/business"
	cartstorage "OpenMarket/module/cart/storage"

	"github.com/gin-gonic/gin"
)

type ItemRequest struct {
	VariantId int `json:"variant_id" binding:"required"`
	Quantity  int `json:"quantity" binding:"required"`
}

func AddToCart(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := appCtx.GetMainDBConnection()
		u, ok := context.MustGet(common.CurrentUser).(common.Requester)
		if !ok {
			panic("cannot get current user")
		}
		userId := u.GetUserId()
		var req ItemRequest
		if err := context.ShouldBindJSON(&req); err != nil {
			panic(common.InvalidRequestError(err))
		}
		storage := cartstorage.NewSQLStore(db)
		biz := cartbusiness.NewAddItemToCartBusiness(storage)
		if err := biz.AddItemToCart(context.Request.Context(), userId, req.VariantId, req.Quantity); err != nil {
			panic(err)
		}
		context.JSON(200, common.SimpleSuccessResponse(true))
	}
}
