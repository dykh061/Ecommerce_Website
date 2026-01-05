package ginseller

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	sellerbusiness "OpenMarket/module/seller/business"
	sellerrepository "OpenMarket/module/seller/repository"
	sellerstorage "OpenMarket/module/seller/storage"

	"github.com/gin-gonic/gin"
)

func DeleteSeller(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := appCtx.GetMainDBConnection()

		u, ok := context.MustGet(common.CurrentUser).(common.Requester)
		if !ok {
			panic(common.ErrUnauthorized(nil))
		}
		userId := u.GetUserId()
		storage := sellerstorage.NewSQLStore(db)
		repo := sellerrepository.NewDeleteSellerRepo(storage)
		finder := sellerrepository.NewGetSellerRepo(storage)
		biz := sellerbusiness.NewDeleteSellerBusiness(repo, finder)
		if err := biz.DeleteSeller(context.Request.Context(), userId); err != nil {
			panic(err)
		}
		context.JSON(200, common.SimpleSuccessResponse(true))
	}
}
