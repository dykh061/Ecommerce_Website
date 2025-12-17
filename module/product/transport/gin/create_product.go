package productgin

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	productbusiness "OpenMarket/module/product/business"
	productmodel "OpenMarket/module/product/model"
	productstorage "OpenMarket/module/product/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProduct(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data productmodel.ProductCreate
		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		store := productstorage.NewSQLStore(db)
		biz := productbusiness.NewCreateProductBusiness(store)
		if err := biz.CreateProduct(ctx.Request.Context(), &data); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data.ID))
	}
}
