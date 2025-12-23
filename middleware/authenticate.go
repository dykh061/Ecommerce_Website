package middleware

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	"OpenMarket/component/tokenprovider/jwt"
	userstorage "OpenMarket/module/user/storage"
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func WrongErrorAuthenHeader(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("wrong authen header "),
		fmt.Sprintf("ErrWrongAuthHeader"),
	)
}

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")
	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", WrongErrorAuthenHeader(nil)
	}
	return parts[1], nil
}

func RequiredAuthenHeader(appCtx appctx.AppContext) func(c *gin.Context) {
	tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())
	return func(c *gin.Context) {
		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))
		if err != nil {
			panic(err)
		}
		db := appCtx.GetMainDBConnection()
		storage := userstorage.NewSQLStore(db)
		payload, err := tokenProvider.Validate(token)
		if err != nil {
			panic(err)
		}
		user, err := storage.FindDataWithCondition(c.Request.Context(), map[string]interface{}{
			"id": payload.UserId,
		})
		if err != nil {
			panic(common.ErrorDB(err))
		}
		if user == nil {
			panic(common.ErrUnauthorized(errors.New("user not found")))
		}
		if user.Status == 0 {
			panic(common.ErrPermission("user has been deleted or banned", errors.New("user has been deleted or banned")))
		}
		c.Set(common.CurrentUser, user)
		c.Next()
	}
}
