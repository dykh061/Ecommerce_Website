package ginseller

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	sellerbusiness "OpenMarket/module/seller/business"
	sellerrepository "OpenMarket/module/seller/repository"
	sellerstorage "OpenMarket/module/seller/storage"

	"github.com/gin-gonic/gin"
)

func GetMyShop(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := appCtx.GetMainDBConnection()
		u, ok := context.MustGet(common.CurrentUser).(common.Requester)
		if !ok {
			panic(common.ErrUnauthorized(nil))
		}
		userId := u.GetUserId()
		storage := sellerstorage.NewSQLStore(db)
		repo := sellerrepository.NewGetSellerRepo(storage)
		biz := sellerbusiness.NewGetMyShopBusiness(repo)
		result, err := biz.GetMyShop(context.Request.Context(), userId)
		if err != nil {
			panic(err)
		}
		result.Mask()
		context.JSON(200, common.SimpleSuccessResponse(result))
	}
}
