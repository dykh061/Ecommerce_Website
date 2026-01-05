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

func UpdateSeller(ctx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := ctx.GetMainDBConnection()
		var data sellermodel.SellerUpdate
		if err := context.ShouldBind(&data); err != nil {
			panic(err)
		}
		requester, ok := context.MustGet(common.CurrentUser).(common.Requester)
		if !ok {
			panic(common.ErrUnauthorized(nil))
		}
		userId := requester.GetUserId()
		store := sellerstorage.NewSQLStore(db)
		repo := sellerrepository.NewUpdateSellerRepo(store)
		finder := sellerrepository.NewGetSellerRepo(store)
		biz := sellerbusiness.NewUpdateSellerBusiness(finder, repo)
		if err := biz.UpdateSeller(context.Request.Context(), userId, data); err != nil {
			panic(err)
		}
		context.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
