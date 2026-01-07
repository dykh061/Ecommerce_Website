package productgin

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	productbusiness "OpenMarket/module/product/business"
	productrepository "OpenMarket/module/product/repository"
	productstorage "OpenMarket/module/product/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

type adjustStockRequest struct {
	VariantId int `json:"variant_id"`
	By        int `json:"by"`
}

func AdjustStock(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var req adjustStockRequest
		if err := c.ShouldBind(&req); err != nil {
			panic(err)
		}
		storage := productstorage.NewSQLStore(db)
		repo := productrepository.NewAdjustStockRepo(storage)
		biz := productbusiness.NewAdjustStockBusiness(repo)
		if err := biz.AdjustStock(c.Request.Context(), req.VariantId, req.By); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
