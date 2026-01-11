package ordergin

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	orderbusiness "OpenMarket/module/order/business"
	orderstorage "OpenMarket/module/order/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

type CreateOrderItemReq struct {
	VariantID int             `json:"variant_id"`
	Quantity  int             `json:"quantity"`
	Price     decimal.Decimal `json:"price"`
}

type CreateOrderReq struct {
	Items []CreateOrderItemReq `json:"items"`
}

func CreateOrder(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var req CreateOrderReq
		if err := c.ShouldBindJSON(&req); err != nil {
			panic(common.InvalidRequestError(err))
		}
		if len(req.Items) == 0 {
			panic(common.InvalidRequestError(
				common.InvalidRequestError(
					nil,
				),
			))
		}
		userID := c.MustGet(common.CurrentUser).(common.Requester).GetUserId()
		items := make([]orderbusiness.CreateOrderItem, 0, len(req.Items))
		for _, item := range req.Items {
			items = append(items, orderbusiness.CreateOrderItem{
				VariantID: item.VariantID,
				Quantity:  item.Quantity,
			})
		}

		store := orderstorage.NewSQLStore(db)
		biz := orderbusiness.NewCreateOrderBusiness(store)
		if err := biz.CreateOrder(c.Request.Context(), userID, items); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
