package productgin

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	productbusiness "OpenMarket/module/product/business"
	productrepository "OpenMarket/module/product/repository"
	productstorage "OpenMarket/module/product/storage"

	"github.com/gin-gonic/gin"
)

func GetImages(appCtx appctx.AppContext) gin.HandlerFunc {

	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(err)
		}
		productID := int(uid.GetLoacalID())
		storage := productstorage.NewSQLStore(db)
		repo := productrepository.NewGetImagesRepo(storage)
		biz := productbusiness.NewGetImagesBusiness(repo)

		result, err := biz.GetImages(c.Request.Context(), productID)
		if err != nil {
			panic(common.ErrCannotListEntity("Images", err))
		}
		c.JSON(200, common.SimpleSuccessResponse(result))
	}

}
