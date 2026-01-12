package productgin

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	productbusiness "OpenMarket/module/product/business"
	productrepository "OpenMarket/module/product/repository"
	productstorage "OpenMarket/module/product/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetCategoryAttributes handles GET /v1/categories/:category_id/attributes
func GetCategoryAttributes(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		categoryID, err := strconv.Atoi(c.Param("category_id"))
		if err != nil {
			panic(common.InvalidRequestError(err))
		}

		storage := productstorage.NewSQLStore(db)
		repo := productrepository.NewGetCategoryAttributesRepo(storage)
		biz := productbusiness.NewGetCategoryAttributesBusiness(repo)

		attributes, err := biz.GetCategoryAttributes(c.Request.Context(), categoryID)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(attributes))
	}
}
