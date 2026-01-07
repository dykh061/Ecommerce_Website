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
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateProduct(appCtx appctx.AppContext) gin.HandlerFunc {
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

		var data productmodel.ProductUpdate
		if err := context.ShouldBind(&data); err != nil {
			panic(err)
		}

		storage := productstorage.NewSQLStore(db)
		sellerstorage := sellerstorage.NewSQLStore(db)
		sfinder := sellerrepository.NewGetSellerRepo(sellerstorage)
		pfinder := productrepository.NewFindProductRepo(storage)
		repo := productrepository.NewUpdateProductRepo(storage)
		biz := productbusiness.NewUpdateProductBusiness(repo, sfinder, pfinder)
		if err := biz.UpdateProduct(context.Request.Context(), userID, productID, &data); err != nil {
			panic(common.ErrCannotUpdateEntity(productmodel.EntityName, err))
		}
		context.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
