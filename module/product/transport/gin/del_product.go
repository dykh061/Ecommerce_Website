package productgin

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	productbusiness "OpenMarket/module/product/business"
	productmodel "OpenMarket/module/product/model"
	productrepository "OpenMarket/module/product/repository"
	productstorage "OpenMarket/module/product/storage"
	sellerrepository "OpenMarket/module/seller/repository"
	sellerstorage "OpenMarket/module/seller/storage"

	"github.com/gin-gonic/gin"
)

func DeleteProduct(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := appCtx.GetMainDBConnection()

		u, ok := context.MustGet(common.CurrentUser).(common.Requester)
		if !ok {
			panic(common.ErrUnauthorized(nil))
		}
		userID := u.GetUserId()
		uid, err := common.FromBase58(context.Param("id"))
		if err != nil {
			panic(err)
		}
		productID := int(uid.GetLoacalID())
		storage := productstorage.NewSQLStore(db)
		sellerstorage := sellerstorage.NewSQLStore(db)
		sfinder := sellerrepository.NewGetSellerRepo(sellerstorage)
		pfinder := productrepository.NewFindProductRepo(storage)
		repo := productrepository.NewDeleteProductRepo(storage)
		biz := productbusiness.NewDeleteProductBusiness(repo, sfinder, pfinder)
		if err := biz.DeleteProduct(context.Request.Context(), userID, productID); err != nil {
			panic(common.ErrCannotDeleteEntity(productmodel.EntityName, err))
		}
		context.JSON(200, common.SimpleSuccessResponse(true))
	}
}
