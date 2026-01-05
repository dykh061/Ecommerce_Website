package productgin

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	productbusiness "OpenMarket/module/product/business"
	productmodel "OpenMarket/module/product/model"
	productrepository "OpenMarket/module/product/repository"
	productstorage "OpenMarket/module/product/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

func ListPublicProduct(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var req common.ListProductRequest
		if err := c.ShouldBindQuery(&req); err != nil {
			panic(err)
		}
		paging := common.Paging{
			Page:  req.Page,
			Limit: req.Limit,
		}

		var sellerID *int
		if req.SellerID != nil {
			uid, err := common.FromBase58(*req.SellerID)
			if err != nil {
				panic(err)
			}
			id := int(uid.GetLoacalID())
			sellerID = &id
		}
		paging.Fulfill()
		filter := &productmodel.Filter{
			Search:     req.Search,
			CategoryID: req.CategoryID,
			SellerID:   sellerID,
		}

		if req.MinPrice != nil {
			v := decimal.NewFromInt(int64(*req.MinPrice))
			filter.MinPrice = &v
		}
		if req.MaxPrice != nil {
			v := decimal.NewFromInt(int64(*req.MaxPrice))
			filter.MaxPrice = &v
		}

		storage := productstorage.NewSQLStore(db)
		repo := productrepository.NewListPublicProductRepo(storage)
		biz := productbusiness.NewListPublicProductBusiness(repo)
		result, err := biz.ListPublicProducts(c.Request.Context(), filter, &paging)
		if err != nil {
			panic(common.ErrCannotListEntity(productmodel.EntityName, err))
		}
		for i := range result {
			result[i].Mask()
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
