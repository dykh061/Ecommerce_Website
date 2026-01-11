package ordergin

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	orderbusiness "OpenMarket/module/order/business"
	orderstorage "OpenMarket/module/order/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CancelOrderRequest struct {
	Id     int    `json:"order_id" binding:"required"`
	Reason string `json:"reason" binding:"required"`
}

func CanCelOrder(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var req CancelOrderRequest
		if err := c.ShouldBind(&req); err != nil {
			panic(err)
		}
		store := orderstorage.NewSQLStore(db)
		biz := orderbusiness.NewCancelOrderBusiness(store)
		if err := biz.CancelOrder(c.Request.Context(), req.Id, req.Reason); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))

	}
}
