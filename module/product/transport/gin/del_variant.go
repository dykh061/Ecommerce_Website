package productgin

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	productbusiness "OpenMarket/module/product/business"
	productrepository "OpenMarket/module/product/repository"
	productstorage "OpenMarket/module/product/storage"
	sellerrepository "OpenMarket/module/seller/repository"
	sellerstorage "OpenMarket/module/seller/storage"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteVariant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := appCtx.GetMainDBConnection()

		u, ok := context.MustGet(common.CurrentUser).(common.Requester)
		if !ok {
			panic(common.ErrUnauthorized(nil))
		}
		userID := u.GetUserId()
		vid, err := strconv.Atoi(context.Param("vid"))
		if err != nil {
			panic(err)
		}
		pid, err := common.FromBase58(context.Param("id"))
		if err != nil {
			panic(err)
		}
		productId := int(pid.GetLoacalID())
		storage := productstorage.NewSQLStore(db)
		sellerstorage := sellerstorage.NewSQLStore(db)
		sfinder := sellerrepository.NewGetSellerRepo(sellerstorage)
		vfinder := productrepository.NewFindProductRepo(storage)
		repo := productrepository.NewDeleteVariantRepo(storage)
		biz := productbusiness.NewDeleteVariantBusiness(repo, sfinder, vfinder)
		if err := biz.DeleteVariant(context.Request.Context(), userID, productId, vid); err != nil {
			panic(common.ErrCannotDeleteEntity("Variant", err))
		}
		context.JSON(200, common.SimpleSuccessResponse(true))
	}
}
