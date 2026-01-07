package ginuser

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	userbusiness "OpenMarket/module/user/business"
	usermodel "OpenMarket/module/user/model"
	srrepository "OpenMarket/module/user/repository"
	userstorage "OpenMarket/module/user/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateUser(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data usermodel.UserUpdate
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}
		u, ok := c.MustGet(common.CurrentUser).(common.Requester)
		if !ok {
			panic(common.ErrUnauthorized(nil))
		}
		storage := userstorage.NewSQLStore(db)

		repo := srrepository.NewUpdateUserRepo(storage)
		finder := srrepository.NewFindUserRepo(storage)

		biz := userbusiness.NewUpdateUserBusiness(repo, finder)

		if err := biz.UpdateUser(c.Request.Context(), u.GetUserId(), data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
