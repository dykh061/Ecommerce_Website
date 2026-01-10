package ginuser

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	userbusiness "OpenMarket/module/user/business"
	userstorage "OpenMarket/module/user/storage"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteAddress(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := appCtx.GetMainDBConnection()
		u, ok := context.MustGet(common.CurrentUser).(common.Requester)
		if !ok {
			panic("cannot find user")
		}
		userId := u.GetUserId()
		id, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			panic(common.InvalidRequestError(err))
		}
		store := userstorage.NewSQLStore(db)
		biz := userbusiness.NewDeleteAddressBusiness(store)
		if err := biz.DeleteAddress(context.Request.Context(), id, userId); err != nil {
			panic(err)
		}
		context.JSON(200, common.SimpleSuccessResponse(true))
	}
}
