package cartgin

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	cartbusiness "OpenMarket/module/cart/business"
	cartrepository "OpenMarket/module/cart/repository"
	cartstorage "OpenMarket/module/cart/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCart(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := appCtx.GetMainDBConnection()
		u, ok := context.MustGet(common.CurrentUser).(common.Requester)
		if !ok {
			panic("no user in context")
		}
		userId := u.GetUserId()
		storage := cartstorage.NewSQLStore(db)
		repo := cartrepository.NewCreateCartRepo(storage)
		find := cartrepository.NewFindCartRepo(storage)
		biz := cartbusiness.NewCreateCartBusiness(repo, find)
		if err := biz.CreateCart(context.Request.Context(), userId); err != nil {
			panic(err)
		}
		context.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
