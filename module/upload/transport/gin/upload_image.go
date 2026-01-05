package uploadgin

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	uploadbusiness "OpenMarket/module/upload/business"
	"io"

	"github.com/gin-gonic/gin"
)

// UpLoadImage là Gin Handler xử lý API upload ảnh
// Nhiệm vụ của handler:
// 1. Nhận file từ request
// 2. Đọc file thành []byte
// 3. Gọi business upload (không upload trực tiếp)
// 4. Trả kết quả cho client
//
// ❗ Handler KHÔNG:
// - Biết S3 / MinIO là gì
// - Biết lưu DB ra sao
// - Chứa logic nghiệp vụ
func UpLoadImage(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		// Lấy file từ multipart/form-data
		// FE phải gửi: form-data key = "file"
		fileHeader, err := c.FormFile("file")
		if err != nil {
			panic(err)
		}

		// Lấy tên folder upload từ form
		// Nếu FE không truyền → mặc định là "img"
		// Ví dụ:
		// folder = "product"
		// folder = "avatar"
		folder := c.DefaultPostForm("folder", "img")

		// Mở file từ FileHeader
		// FileHeader chỉ là metadata, muốn đọc phải Open()
		file, err := fileHeader.Open()
		if err != nil {
			panic(err)
		}
		// Đảm bảo file được đóng sau khi xử lý xong
		defer file.Close()
		// Đọc toàn bộ nội dung file vào bộ nhớ
		// Kết quả là []byte để truyền cho tầng business

		dataBytes, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}

		// Khởi tạo upload business
		// UploadProvider được inject từ AppContext
		// → có thể là:
		// - S3 / MinIO
		// - Local storage
		// - Cloud khác
		biz := uploadbusiness.NewUploadBusiness(appCtx.UploadProvider())

		// Gọi business xử lý upload
		// Business sẽ:
		// - Validate file (nếu có)
		// - Sinh đường dẫn
		// - Gọi UploadProvider.SaveFileUploaded()
		// - Trả về thông tin Image (URL)
		img, err := biz.UploadFile(c.Request.Context(), dataBytes, folder, fileHeader.Filename)
		if err != nil {
			c.JSON(400, common.NewCustomError(err, "invalid image file", "ErrInvalidImage"))
			return
		}
		c.JSON(200, common.SimpleSuccessResponse(img))
	}
}
