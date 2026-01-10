package ginuser

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	userbusiness "OpenMarket/module/user/business"
	usermodel "OpenMarket/module/user/model"
	userstorage "OpenMarket/module/user/storage"

	"github.com/gin-gonic/gin"
)

func GetListAddress(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		u, ok := c.MustGet(common.CurrentUser).(common.Requester)
		if !ok {
			panic("cannot find user")
		}

		userId := u.GetUserId()
		var result []usermodel.UserAddress
		store := userstorage.NewSQLStore(db)
		biz := userbusiness.NewGetListAddressBusiness(store)
		result, err := biz.GetListAddress(c.Request.Context(), userId)
		if err != nil {
			panic(err)
		}
		for v := range result {
			result[v].Mask()
		}
		c.JSON(200, common.NewSuccessResponse(result, nil, nil))

	}
}
