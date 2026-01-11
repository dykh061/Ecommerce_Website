package ordergin

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	orderbusiness "OpenMarket/module/order/business"
	orderstorage "OpenMarket/module/order/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderDetailReq struct {
	OrderId int `uri:"order_id" binding:"required"`
}

func GetDetailOrder(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var req OrderDetailReq
		if err := context.ShouldBindUri(&req); err != nil {
			panic(err)
		}
		store := orderstorage.NewSQLStore(db)
		biz := orderbusiness.NewGetDetailOrderBusiness(store)
		result, err := biz.GetDetailOrder(context.Request.Context(), req.OrderId)
		if err != nil {
			panic(err)
		}
		context.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
