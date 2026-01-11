package ordergin

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	orderbusiness "OpenMarket/module/order/business"
	ordermodel "OpenMarket/module/order/model"
	orderstorage "OpenMarket/module/order/storage"
	productmodel "OpenMarket/module/product/model"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

func GetListOrder(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var req common.GetListOrderRequest
		if err := context.ShouldBindQuery(&req); err != nil {
			panic(err)
		}
		u, ok := context.MustGet(common.CurrentUser).(common.Requester)
		if !ok {
			panic(common.ErrPermission("no permission", nil))
		}
		paging := common.Paging{
			Page:  req.Page,
			Limit: req.Limit,
		}
		paging.Fulfill()
		filter := &ordermodel.FilterOrder{
			Status: req.Status,
			UserId: int(u.GetUserId()),
		}
		if req.MinPrice != nil {
			v := decimal.NewFromInt(int64(*req.MinPrice))
			filter.MinPrice = &v
		}
		if req.MaxPrice != nil {
			v := decimal.NewFromInt(int64(*req.MaxPrice))
			filter.MaxPrice = &v
		}
		store := orderstorage.NewSQLStore(db)
		biz := orderbusiness.NewGetListOrderBiz(store)
		result, err := biz.GetListOrder(context.Request.Context(), filter, &paging)
		if err != nil {
			panic(common.ErrCannotListEntity(productmodel.EntityName, err))
		}
		context.JSON(200, common.NewSuccessResponse(result, paging, filter))
	}
}
