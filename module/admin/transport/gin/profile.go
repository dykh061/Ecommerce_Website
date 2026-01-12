package admingin

import (
	"OpenMarket/common"
	"OpenMarket/middleware"
	adminmodel "OpenMarket/module/admin/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Profile returns the authenticated staff info
// GET /v1/admin/profile
func Profile() gin.HandlerFunc {
	return func(c *gin.Context) {
		staffRaw, exists := c.Get(middleware.CurrentStaff)
		if !exists {
			panic(common.ErrUnauthorized(nil))
		}

		staff, ok := staffRaw.(*adminmodel.Staff)
		if !ok {
			panic(common.ErrInternal(nil))
		}

		staff.Mask()
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(staff))
	}
}
