package productgin

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	productbusiness "OpenMarket/module/product/business"
	productrepository "OpenMarket/module/product/repository"
	productstorage "OpenMarket/module/product/storage"

	"github.com/gin-gonic/gin"
)

func ListVariant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(err)
		}
		productID := int(uid.GetLoacalID())
		storage := productstorage.NewSQLStore(db)
		repo := productrepository.NewListVariantRepo(storage)
		biz := productbusiness.NewListVariantBusiness(repo)
		result, err := biz.ListVariant(c.Request.Context(), productID)
		if err != nil {
			panic(common.ErrCannotListEntity("Variant", err))
		}
		c.JSON(200, common.SimpleSuccessResponse(result))
	}
}
