package ginseller

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	sellerbusiness "OpenMarket/module/seller/business"
	sellermodel "OpenMarket/module/seller/model"
	sellerrepository "OpenMarket/module/seller/repository"
	sellerstorage "OpenMarket/module/seller/storage"

	"github.com/gin-gonic/gin"
)

func ListSeller(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var filter sellermodel.SellerFilter
		if err := context.ShouldBindQuery(&filter); err != nil {
			panic(common.ErrCannotListEntity(sellermodel.EntityName, err))
		}

		var pagingData common.Paging
		if err := context.ShouldBindQuery(&pagingData); err != nil {
			panic(common.InvalidRequestError(err))
		}

		pagingData.Fulfill()

		storage := sellerstorage.NewSQLStore(db)
		repo := sellerrepository.NewListSellerRepo(storage)
		biz := sellerbusiness.NewListSellerBusiness(repo)

		results, err := biz.ListSellers(context.Request.Context(), &filter, &pagingData)
		if err != nil {
			panic(common.ErrCannotListEntity(sellermodel.EntityName, err))
		}

		for i := range results {
			results[i].Mask()
		}
		context.JSON(200, common.NewSuccessResponse(results, pagingData, filter))
	}
}
