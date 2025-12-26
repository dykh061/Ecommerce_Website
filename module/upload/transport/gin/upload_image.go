package uploadgin

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	uploadbusiness "OpenMarket/module/upload/business"
	"io"

	"github.com/gin-gonic/gin"
)

func UpLoadImage(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")
		if err != nil {
			panic(err)
		}

		folder := c.DefaultPostForm("folder", "img")

		file, err := fileHeader.Open()
		if err != nil {
			panic(err)
		}
		defer file.Close()

		dataBytes, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}

		biz := uploadbusiness.NewUploadBusiness(appCtx.UploadProvider())
		img, err := biz.UploadFile(c.Request.Context(), dataBytes, folder, fileHeader.Filename)
		if err != nil {
			c.JSON(400, common.NewCustomError(err, "invalid image file", "ErrInvalidImage"))
			return
		}
		c.JSON(200, common.SimpleSuccessResponse(img))
	}
}
