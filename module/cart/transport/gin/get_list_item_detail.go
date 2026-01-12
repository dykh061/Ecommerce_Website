package cartgin

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	cartbusiness "OpenMarket/module/cart/business"
	cartstorage "OpenMarket/module/cart/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetListItemDetail(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := appCtx.GetMainDBConnection()
		u, ok := context.MustGet(common.CurrentUser).(common.Requester)
		if !ok {
			panic("cannot get current user")
		}
		userId := u.GetUserId()

		storage := cartstorage.NewSQLStore(db)
		biz := cartbusiness.NewListCartItemDetailBusiness(storage)
		result, err := biz.ListCartItemDetail(context.Request.Context(), userId)
		if err != nil {
			panic(err)
		}
		context.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
