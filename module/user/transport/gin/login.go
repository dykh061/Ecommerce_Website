package ginuser

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	"OpenMarket/component/hasher"
	"OpenMarket/component/tokenprovider/jwt"
	userbusiness "OpenMarket/module/user/business"
	usermodel "OpenMarket/module/user/model"
	srrepository "OpenMarket/module/user/repository"
	userstorage "OpenMarket/module/user/storage"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData usermodel.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(common.InvalidRequestError(err))
		}

		db := appCtx.GetMainDBConnection()

		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

		storage := userstorage.NewSQLStore(db)
		hasher := hasher.NewBcryptHasher(bcrypt.DefaultCost)
		frepo := srrepository.NewFindUserWithEmailRepo(storage)
		biz := userbusiness.NewLoginBusiness(frepo, tokenProvider, hasher, 60*60*24*30)
		account, err := biz.Login(c.Request.Context(), &loginUserData)
		if err != nil {
			panic(err)
		}
		c.JSON(200, common.SimpleSuccessResponse(account))

	}
}
