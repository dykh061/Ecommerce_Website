package ginseller

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	sellerbusiness "OpenMarket/module/seller/business"
	sellermodel "OpenMarket/module/seller/model"
	sellerstorage "OpenMarket/module/seller/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateSeller(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data sellermodel.SellerCreate
		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		store := sellerstorage.NewSQLStore(db)
		biz := sellerbusiness.NewCreateSellerBusiness(store)

		if err := biz.CreateSeller(ctx.Request.Context(), &data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
