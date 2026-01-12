package admingin

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	"OpenMarket/component/hasher"
	"OpenMarket/component/tokenprovider/jwt"
	adminbusiness "OpenMarket/module/admin/business"
	adminmodel "OpenMarket/module/admin/model"
	adminstorage "OpenMarket/module/admin/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login handles staff authentication
// POST /v1/admin/auth/login
func Login(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data adminmodel.StaffLogin

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.InvalidRequestError(err))
		}

		db := appCtx.GetMainDBConnection()
		store := adminstorage.NewSQLStore(db)
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())
		bcryptHasher := hasher.NewBcryptHasher(12)

		biz := adminbusiness.NewLoginStaffBusiness(
			store,
			tokenProvider,
			bcryptHasher,
			60*60*24*7, // 7 days expiry for admin token
		)

		token, staff, err := biz.Login(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(gin.H{
			"token": token,
			"staff": staff,
		}))
	}
}
