package ginuser

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	userbusiness "OpenMarket/module/user/business"
	userstorage "OpenMarket/module/user/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteUser(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		db := appCtx.GetMainDBConnection()
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(err)
		}
		store := userstorage.NewSQLStore(db)
		biz := userbusiness.NewDeleteUserBusiness(store)
		if err := biz.DeleteUser(c.Request.Context(), id); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
