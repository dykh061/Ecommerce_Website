package ginuser

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	userbusiness "OpenMarket/module/user/business"
	usermodel "OpenMarket/module/user/model"
	userstorage "OpenMarket/module/user/storage"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateAddress(appCtx appctx.AppContext) gin.HandlerFunc {
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
		var req usermodel.UserAddressUpdate
		if err := context.ShouldBind(&req); err != nil {
			panic(common.InvalidRequestError(err))
		}
		store := userstorage.NewSQLStore(db)
		biz := userbusiness.NewUpdateAddressBusiness(store)
		if err := biz.UpdateAddress(context.Request.Context(), id, userId, &req); err != nil {
			panic(err)
		}
		context.JSON(200, common.SimpleSuccessResponse(true))
	}
}
