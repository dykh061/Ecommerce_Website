package ginuser

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	userbusiness "OpenMarket/module/user/business"
	srrepository "OpenMarket/module/user/repository"
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

		u, ok := c.MustGet(common.CurrentUser).(common.Requester)
		if !ok {
			panic(common.ErrUnauthorized(nil))
		}

		if u.GetUserId() != id {
			panic(common.ErrPermission("you don't have permission to delete this user", nil))
		}
		store := userstorage.NewSQLStore(db)
		frepo := srrepository.NewFindUserRepo(store)
		drepo := srrepository.NewDeleteUserRepo(store)
		biz := userbusiness.NewDeleteUserBusiness(frepo, drepo)
		if err := biz.DeleteUser(c.Request.Context(), id); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
