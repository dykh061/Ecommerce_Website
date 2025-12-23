package ginuser

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	userbusiness "OpenMarket/module/user/business"
	usermodel "OpenMarket/module/user/model"
	userstorage "OpenMarket/module/user/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateUser(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data usermodel.UserUpdate
		var id int
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(err)
		}

		storage := userstorage.NewSQLStore(db)

		biz := userbusiness.NewUpdateUserBusiness(storage)

		if err := biz.UpdateUser(c.Request.Context(), id, data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
