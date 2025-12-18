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
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "ID Invalid",
			})
		}

		storage := userstorage.NewSQLStore(db)

		biz := userbusiness.NewUpdateUserBusiness(storage)

		if err := biz.UpdateUser(c, id, data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
