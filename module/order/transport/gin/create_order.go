package ordergin

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	orderbusiness "OpenMarket/module/order/business"
	orderstorage "OpenMarket/module/order/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrder(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		type createOrderReq struct {
			AddressID string `json:"address_id"`
		}

		var req createOrderReq
		if err := c.ShouldBindJSON(&req); err != nil {
			panic(common.InvalidRequestError(err))
		}

		uid, err := common.FromBase58(req.AddressID)
		if err != nil {
			panic(common.InvalidRequestError(err))
		}

		addressId := int(uid.GetLoacalID())
		if addressId <= 0 {
			panic(common.InvalidRequestError(nil))
		}

		userID := c.MustGet(common.CurrentUser).(common.Requester).
			GetUserId()

		store := orderstorage.NewSQLStore(db)
		biz := orderbusiness.NewCreateOrderBusiness(store)

		if err := biz.CreateOrder(
			c.Request.Context(),
			userID,
			addressId,
		); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
