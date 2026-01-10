package ginuser

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	userbusiness "OpenMarket/module/user/business"
	usermodel "OpenMarket/module/user/model"
	userstorage "OpenMarket/module/user/storage"

	"github.com/gin-gonic/gin"
)

func CreateAddress(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		u, ok := c.MustGet(common.CurrentUser).(common.Requester)
		if !ok {
			panic("cannot find user")
		}

		var req usermodel.UserAddressCreate
		if err := c.ShouldBind(&req); err != nil {
			panic(common.InvalidRequestError(err))
		}

		userId := u.GetUserId()
		store := userstorage.NewSQLStore(db)
		biz := userbusiness.NewCreateAddressBusiness(store)
		if err := biz.CreateAddress(c.Request.Context(), userId, &req); err != nil {
			panic(err)
		}

		c.JSON(200, common.SimpleSuccessResponse(true))
	}
}
