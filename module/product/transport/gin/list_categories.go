package productgin

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	productbusiness "OpenMarket/module/product/business"
	productrepository "OpenMarket/module/product/repository"
	productstorage "OpenMarket/module/product/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListCategories handles GET /v1/categories
func ListCategories(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		storage := productstorage.NewSQLStore(db)
		repo := productrepository.NewListAllCategoriesRepo(storage)
		biz := productbusiness.NewListCategoriesBusiness(repo)

		categories, err := biz.ListCategories(c.Request.Context())
		if err != nil {
			panic(common.ErrCannotListEntity("Category", err))
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(categories))
	}
}
