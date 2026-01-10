package cartgin

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	cartbusiness "OpenMarket/module/cart/business"
	cartrepository "OpenMarket/module/cart/repository"
	cartstorage "OpenMarket/module/cart/storage"

	"github.com/gin-gonic/gin"
)

type AdjustItemRequest struct {
	VariantId int `json:"variant_id" binding:"required"`
	Quantity  int `json:"quantity" binding:"required"`
}

func AdJustProduct(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := appCtx.GetMainDBConnection()
		u, ok := context.MustGet(common.CurrentUser).(common.Requester)
		if !ok {
			panic("cannot get current user")
		}
		userId := u.GetUserId()
		var req AdjustItemRequest
		if err := context.ShouldBind(&req); err != nil {
			panic(common.InvalidRequestError(err))
		}
		if req.Quantity == 0 {
			panic(common.InvalidRequestError(nil))
		}
		storage := cartstorage.NewSQLStore(db)
		finder := cartrepository.NewFindCartRepo(storage)
		findItem := cartrepository.NewFindCartItemRepo(storage)
		repo := cartrepository.NewUpdateCartItemRepo(storage)
		biz := cartbusiness.NewAdjustProductInCartBusiness(repo, finder, findItem)
		if err := biz.AdjustProductInCart(context.Request.Context(), req.Quantity, userId, req.VariantId); err != nil {
			panic(err)
		}
		context.JSON(200, common.SimpleSuccessResponse(true))
	}
}
