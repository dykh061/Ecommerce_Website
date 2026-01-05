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

func CreateProduct(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := appCtx.GetMainDBConnection()
		requester, ok := ctx.MustGet(common.CurrentUser).(common.Requester)
		if !ok {
			panic(common.ErrUnauthorized(nil))
		}
		var data productmodel.ProductCreate
		if err := ctx.ShouldBind(&data); err != nil {
			panic(common.InvalidRequestError(nil))
		}
		productStore := productstorage.NewSQLStore(db)
		sellerStore := sellerstorage.NewSQLStore(db)
		productRepo := productrepository.NewCreateProductRepo(productStore)
		sellerRepo := sellerrepository.NewGetSellerRepo(sellerStore)
		biz := productbusiness.NewCreateProductBusiness(productRepo, sellerRepo)
		if err := biz.CreateProduct(ctx.Request.Context(), requester.GetUserId(), &data); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
