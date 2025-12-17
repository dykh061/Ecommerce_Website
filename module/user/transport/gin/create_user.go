package ginuser

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	userbusiness "OpenMarket/module/user/business"
	usermodel "OpenMarket/module/user/model"
	userstorage "OpenMarket/module/user/storage"

	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUser là hàm xử lý HTTP request tạo user.
//
// Hàm này sẽ:
// 1. Nhận request từ client (Gin)
// 2. Gọi business để thực hiện nghiệp vụ tạo user
// 3. Trả về response cho client (Gin)
//
// Hàm này KHÔNG làm việc trực tiếp với database.
// Nó chỉ điều phối luồng dữ liệu giữa client và business.
func CreateUser(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data usermodel.UserCreate
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := userstorage.NewSQLStore(db)
		biz := userbusiness.NewCreateUserBusiness(store)

		if err := biz.CreateUser(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.ID))
	}
}
