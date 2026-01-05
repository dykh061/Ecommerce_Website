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
	"github.com/shopspring/decimal"
)

func ListSellerProduct(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		u, ok := c.MustGet(common.CurrentUser).(common.Requester)
		if !ok {
			panic(common.ErrUnauthorized(nil))
		}
		userID := u.GetUserId()

		var req common.ListProductRequest
		if err := c.ShouldBindQuery(&req); err != nil {
			panic(err)
		}

		paging := common.Paging{
			Page:  req.Page,
			Limit: req.Limit,
		}
		paging.Fulfill()
		filter := &productmodel.Filter{
			Search:     req.Search,
			CategoryID: req.CategoryID,
		}
		if req.MinPrice != nil {
			v := decimal.NewFromInt(int64(*req.MinPrice))
			filter.MinPrice = &v
		}
		if req.MaxPrice != nil {
			v := decimal.NewFromInt(int64(*req.MaxPrice))
			filter.MaxPrice = &v
		}

		sellerStore := sellerstorage.NewSQLStore(db)
		sellerFinder := sellerrepository.NewGetSellerRepo(sellerStore)
		store := productstorage.NewSQLStore(db)
		repo := productrepository.NewListPublicProductRepo(store)
		biz := productbusiness.NewListSellerProductBusiness(repo, sellerFinder)

		result, err := biz.ListSellerProducts(c.Request.Context(), userID, filter, &paging)
		if err != nil {
			panic(common.ErrCannotListEntity(productmodel.EntityName, err))
		}

		for i := range result {
			result[i].Mask()
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
