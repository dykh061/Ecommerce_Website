package productgin

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	productbusiness "OpenMarket/module/product/business"
	productrepository "OpenMarket/module/product/repository"
	productstorage "OpenMarket/module/product/storage"
	sellerrepository "OpenMarket/module/seller/repository"
	sellerstorage "OpenMarket/module/seller/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// SetMainGallery handles PATCH /v1/seller/products/:id/galleries/:gallery_id/main
func SetMainGallery(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		u, ok := c.MustGet(common.CurrentUser).(common.Requester)
		if !ok {
			panic(common.ErrUnauthorized(nil))
		}
		userID := u.GetUserId()

		// Parse product ID
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.InvalidRequestError(err))
		}
		productID := int(uid.GetLoacalID())

		// Parse gallery ID
		galleryID, err := strconv.Atoi(c.Param("gallery_id"))
		if err != nil {
			panic(common.InvalidRequestError(err))
		}

		storage := productstorage.NewSQLStore(db)
		sellerStorage := sellerstorage.NewSQLStore(db)
		sfinder := sellerrepository.NewGetSellerRepo(sellerStorage)
		pfinder := productrepository.NewFindProductRepo(storage)
		galleryRepo := productrepository.NewGalleryManagementRepo(storage)

		biz := productbusiness.NewGalleryManagementBusiness(sfinder, pfinder, galleryRepo)

		if err := biz.SetMainGallery(c.Request.Context(), userID, productID, galleryID); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

// DeleteGallery handles DELETE /v1/seller/products/:id/galleries/:gallery_id
func DeleteGallery(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		u, ok := c.MustGet(common.CurrentUser).(common.Requester)
		if !ok {
			panic(common.ErrUnauthorized(nil))
		}
		userID := u.GetUserId()

		// Parse product ID
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.InvalidRequestError(err))
		}
		productID := int(uid.GetLoacalID())

		// Parse gallery ID
		galleryID, err := strconv.Atoi(c.Param("gallery_id"))
		if err != nil {
			panic(common.InvalidRequestError(err))
		}

		storage := productstorage.NewSQLStore(db)
		sellerStorage := sellerstorage.NewSQLStore(db)
		sfinder := sellerrepository.NewGetSellerRepo(sellerStorage)
		pfinder := productrepository.NewFindProductRepo(storage)
		galleryRepo := productrepository.NewGalleryManagementRepo(storage)

		biz := productbusiness.NewGalleryManagementBusiness(sfinder, pfinder, galleryRepo)

		if err := biz.DeleteGallery(c.Request.Context(), userID, productID, galleryID); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
