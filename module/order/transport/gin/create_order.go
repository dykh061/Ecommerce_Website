package ordergin

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	orderbusiness "OpenMarket/module/order/business"
	orderstorage "OpenMarket/module/order/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateOrderReq struct {
	AddressId int `json:"address_id"`
}

func CreateOrder(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var req CreateOrderReq
		if err := c.ShouldBindJSON(&req); err != nil {
			panic(common.InvalidRequestError(err))
		}

		userID := c.MustGet(common.CurrentUser).(common.Requester).GetUserId()

		store := orderstorage.NewSQLStore(db)
		biz := orderbusiness.NewCreateOrderBusiness(store)
		if err := biz.CreateOrder(c.Request.Context(), userID, req.AddressId); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
