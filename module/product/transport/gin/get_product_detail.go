package productgin

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	productbusiness "OpenMarket/module/product/business"
	productrepository "OpenMarket/module/product/repository"
	productstorage "OpenMarket/module/product/storage"

	"github.com/gin-gonic/gin"
)

func GetProductDetail(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := appCtx.GetMainDBConnection()
		uid, err := common.FromBase58(context.Param("id"))
		if err != nil {
			panic(err)
		}
		productID := int(uid.GetLoacalID())
		storage := productstorage.NewSQLStore(db)
		finder := productrepository.NewGetProductRepo(storage)
		listVariant := productrepository.NewListVariantRepo(storage)
		getImages := productrepository.NewGetImagesRepo(storage)
		biz := productbusiness.NewGetProductDetailBusiness(finder, listVariant, getImages)
		result, err := biz.GetProductDetail(context.Request.Context(), productID)
		if err != nil {
			panic(common.ErrCannotListEntity("Product", err))
		}
		context.JSON(200, common.SimpleSuccessResponse(result))
	}
}
