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

func FindUser(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		store := userstorage.NewSQLStore(db)
		biz := userbusiness.NewFindUserBusiness(store)

		user, err := biz.FindUser(c.Request.Context(), map[string]interface{}{"id": id})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(user))
	}
}
