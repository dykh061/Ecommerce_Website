package productgin

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	productbusiness "OpenMarket/module/product/business"
	productrepository "OpenMarket/module/product/repository"
	productstorage "OpenMarket/module/product/storage"
	sellerrepository "OpenMarket/module/seller/repository"
	sellerstorage "OpenMarket/module/seller/storage"
	uploadbusiness "OpenMarket/module/upload/business"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateProductGallery(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		db := appCtx.GetMainDBConnection()
		u, ok := c.MustGet(common.CurrentUser).(common.Requester)
		if !ok {
			panic(common.ErrUnauthorized(nil))
		}
		userID := u.GetUserId()

		// 1. Lấy productId từ URL
		productId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(err)
		}

		// 2. Lấy file
		fileHeader, err := c.FormFile("file")
		if err != nil {
			panic(err)
		}

		file, err := fileHeader.Open()
		if err != nil {
			panic(err)
		}
		defer file.Close()

		dataBytes, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}

		// is_main (optional) - robust parse
		isMainStr := c.DefaultPostForm("is_main", "false")
		isMain, err := strconv.ParseBool(isMainStr)
		uploadBiz := uploadbusiness.NewUploadBusiness(appCtx.UploadProvider())
		storage := productstorage.NewSQLStore(db)
		sellerStorage := sellerstorage.NewSQLStore(db)
		galleryRepo := productrepository.NewCreateGalleryRepo(storage)
		sfinder := sellerrepository.NewGetSellerRepo(sellerStorage)
		pfinder := productrepository.NewFindProductRepo(storage)
		createGalleryBiz := productbusiness.NewCreateGalleryBiz(uploadBiz, galleryRepo, sfinder, pfinder)
		// 4. Gọi business
		if err := createGalleryBiz.CreateProductGallery(
			c.Request.Context(),
			productId,
			userID,
			dataBytes,
			fileHeader.Filename,
			isMain,
		); err != nil {
			panic(err)
		}

		// 5. Trả kết quả
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
