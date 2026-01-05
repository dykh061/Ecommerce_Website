package uploadgin

import (
	"OpenMarket/common"
	uploadprovider "OpenMarket/component/uploadProvider"

	"bytes"
	"context"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"path/filepath"
	"time"
)

// uploadBusiness là tầng nghiệp vụ upload
// Nó KHÔNG:
// - Biết HTTP
// - Biết Gin
// - Biết request đến từ đâu
//
// Nó CHỈ:
// - Validate file
// - Lấy metadata ảnh
// - Gọi UploadProvider để lưu file
type uploadBusiness struct {
	provider uploadprovider.UploadProvider
}

// NewUploadBusiness khởi tạo uploadBusiness
// UploadProvider được inject từ ngoài vào
// → đảm bảo đảo chiều phụ thuộc (DIP)
func NewUploadBusiness(provider uploadprovider.UploadProvider) *uploadBusiness {
	return &uploadBusiness{provider: provider}
}

// UploadFile xử lý nghiệp vụ upload ảnh
//
// Input:
// - ctx: context request
// - data: nội dung file dạng []byte
// - folder: thư mục lưu (img, product, avatar...)
// - filename: tên file gốc từ client
//
// Output:
// - *common.Image: chứa url + metadata
func (b *uploadBusiness) UploadFile(ctx context.Context, data []byte, folder, filename string) (*common.Image, error) {
	// Tạo buffer từ []byte
	// Mục đích: truyền vào image.DecodeConfig
	fileBytes := bytes.NewBuffer(data)
	w, h, err := getImageDimension(fileBytes)
	if err != nil {
		return nil, err
	}
	fileExt := filepath.Ext(filename)

	// Đổi tên file để tránh trùng
	// Dùng Nanosecond cho độ unique cao
	filename = fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt)

	// Gọi provider để upload file
	// Business KHÔNG biết:
	// - Lưu local
	// - Hay S3
	// - Hay MinIO
	img, err := b.provider.SaveFileUploaded(ctx, data, fmt.Sprintf("%s/%s", folder, filename))
	if err != nil {
		return nil, err
	}

	img.Width = w
	img.Height = h
	img.Extension = fileExt
	return img, nil
}

// getImageDimension đọc metadata ảnh
// Trả về width, height
// Không decode toàn bộ ảnh → rất nhanh
func getImageDimension(reader io.Reader) (int, int, error) {
	img, _, err := image.DecodeConfig(reader)
	if err != nil {
		log.Println("Failed to decode image:", err)
		return 0, 0, err
	}
	return img.Width, img.Height, nil
}
