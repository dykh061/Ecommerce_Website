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
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateVariant(appCtx appctx.AppContext) gin.HandlerFunc {
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
		var data productmodel.VariantUpdate

		if err := context.ShouldBind(&data); err != nil {
			panic(err)
		}
		if data.Price == nil && data.StockQuantity == nil {
			panic(common.InvalidRequestError(
				errors.New("nothing to update"),
			))
		}

		storage := productstorage.NewSQLStore(db)
		sellerstorage := sellerstorage.NewSQLStore(db)
		sfinder := sellerrepository.NewGetSellerRepo(sellerstorage)
		vfinder := productrepository.NewFindProductRepo(storage)
		repo := productrepository.NewUpdateVariantRepo(storage)
		biz := productbusiness.NewUpdateVariantBusiness(repo, sfinder, vfinder)
		if err := biz.UpdateVariant(context.Request.Context(), userID, productId, vid, &data); err != nil {
			panic(common.ErrCannotUpdateEntity(productmodel.EntityName, err))
		}
		context.JSON(200, common.SimpleSuccessResponse(true))
	}
}
