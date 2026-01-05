package ginuser

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Profile(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		u, ok := c.MustGet(common.CurrentUser).(common.Requester)
		if !ok {
			panic(common.ErrUnauthorized(errors.New("missing auth context")))
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(u))
	}
}
