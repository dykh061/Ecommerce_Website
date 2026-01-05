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
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateVariant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := appCtx.GetMainDBConnection()
		u, ok := context.MustGet(common.CurrentUser).(common.Requester)
		if !ok {
			panic(common.ErrUnauthorized(nil))
		}
		userId := u.GetUserId()
		var data productmodel.VariantCreate
		if err := context.ShouldBind(&data); err != nil {
			panic(common.InvalidRequestError(err))
		}
		pid, err := strconv.Atoi(context.Param("id"))
		//pid, err := common.FromBase58(context.Param("id"))
		if err != nil {
			panic(err)
		}
		storage := productstorage.NewSQLStore(db)
		sellerStorage := sellerstorage.NewSQLStore(db)
		sellerFinder := sellerrepository.NewGetSellerRepo(sellerStorage)
		productFinder := productrepository.NewFindProductRepo(storage)
		repo := productrepository.NewCreateVariantRepo(storage)
		biz := productbusiness.NewCreateVariantBusiness(repo, sellerFinder, productFinder)
		if err := biz.CreateVariant(context.Request.Context(), userId, pid, &data); err != nil {
			panic(common.ErrCannotCreateEntity("Variant", err))
		}
		context.JSON(200, common.SimpleSuccessResponse(true))
	}
}
