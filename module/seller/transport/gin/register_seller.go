package ginseller

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	sellerbusiness "OpenMarket/module/seller/business"
	sellermodel "OpenMarket/module/seller/model"
	sellerrepository "OpenMarket/module/seller/repository"
	sellerstorage "OpenMarket/module/seller/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateSeller(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data sellermodel.SellerCreate
		if err := ctx.ShouldBind(&data); err != nil {
			panic(common.InvalidRequestError(err))
		}
		u, ok := ctx.MustGet(common.CurrentUser).(common.Requester)
		if !ok {
			panic(common.ErrUnauthorized(nil))
		}
		data.UserID = u.GetUserId()
		store := sellerstorage.NewSQLStore(db)
		crepo := sellerrepository.NewCreateSellerRepo(store)
		frepo := sellerrepository.NewFindSellerRepo(store)
		biz := sellerbusiness.NewCreateSellerBusiness(crepo, frepo)

		if err := biz.CreateSeller(ctx.Request.Context(), u.GetUserId(), &data); err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
