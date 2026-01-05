package ginuser

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	"OpenMarket/component/hasher"
	userbusiness "OpenMarket/module/user/business"
	srrepository "OpenMarket/module/user/repository"

	userstorage "OpenMarket/module/user/storage"
	"errors"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=8"`
}

func ChangePassword(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var req ChangePasswordRequest
		if err := context.ShouldBind(&req); err != nil {
			panic(err)
		}
		u, ok := context.MustGet(common.CurrentUser).(common.Requester)
		if !ok {
			panic(common.ErrUnauthorized(errors.New("missing auth context")))
		}

		store := userstorage.NewSQLStore(db)
		hasher := hasher.NewBcryptHasher(bcrypt.DefaultCost)
		changerepo := srrepository.NewChangePasswordRepo(store)
		findrepo := srrepository.NewFindUserRepo(store)
		biz := userbusiness.NewChangePasswordBiz(changerepo, findrepo, hasher)
		if err := biz.ChangePassword(context, u.GetUserId(), req.OldPassword, req.NewPassword); err != nil {
			panic(err)
		}
		context.JSON(200, common.SimpleSuccessResponse(true))
	}
}
