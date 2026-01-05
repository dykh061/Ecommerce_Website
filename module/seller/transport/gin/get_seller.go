package ginseller

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	sellerbusiness "OpenMarket/module/seller/business"
	sellerrepository "OpenMarket/module/seller/repository"
	sellerstorage "OpenMarket/module/seller/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSeller(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := appCtx.GetMainDBConnection()
		//id, err := strconv.Atoi(context.Param("id"))
		uid, err := common.FromBase58(context.Param("id"))
		if err != nil {
			panic(err)
		}
		storage := sellerstorage.NewSQLStore(db)
		repo := sellerrepository.NewGetSellerWithIdRepo(storage)
		biz := sellerbusiness.NewGetSellerBusiness(repo)
		result, err := biz.GetSeller(context.Request.Context(), int(uid.GetLoacalID()))
		if err != nil {
			panic(err)
		}
		result.Mask()
		context.JSON(http.StatusOK, common.SimpleSuccessResponse(result))

	}
}
