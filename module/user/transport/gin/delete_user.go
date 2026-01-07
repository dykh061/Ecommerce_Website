package ginuser

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	userbusiness "OpenMarket/module/user/business"
	srrepository "OpenMarket/module/user/repository"
	userstorage "OpenMarket/module/user/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteUser(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		db := appCtx.GetMainDBConnection()
		u, ok := c.MustGet(common.CurrentUser).(common.Requester)
		if !ok {
			panic(common.ErrUnauthorized(nil))
		}
		store := userstorage.NewSQLStore(db)
		frepo := srrepository.NewFindUserRepo(store)
		drepo := srrepository.NewDeleteUserRepo(store)
		biz := userbusiness.NewDeleteUserBusiness(frepo, drepo)
		if err := biz.DeleteUser(c.Request.Context(), u.GetUserId()); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
