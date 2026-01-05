package uploadprovider

import (
	"OpenMarket/common"
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// s3Provider quản lý việc upload file lên S3-compatible storage
//
// Có thể sử dụng cho:
// - AWS S3
// - MinIO
// - Any S3-compatible service
//
// Vai trò:
// - Nhận dữ liệu file dạng []byte
// - Upload file lên bucket
// - Trả về link public của file
//
// ❗ s3Provider KHÔNG:
// - Biết HTTP / Gin
// - Biết business logic
// - Biết DB

type s3Provider struct {
	bucketName   string           //tên của bucket chứa file
	region       string           // vùng region của bucketS3
	apiKey       string           // apiKey của tài khoản để xác thực
	secret       string           // mật khẩu của tài khoản để xác thực
	endpoint     string           // địa chỉ endpoint của s3
	publicDomain string           // tên miền public để truy cập file
	session      *session.Session // phiên làm việc với s3
}

// NewS3Provider khởi tạo một s3Provider mới
//
// Tại đây:
// - Cấu hình kết nối S3 / MinIO
// - Tạo session dùng lại cho các request upload
//
// ❗ Hàm này thường được gọi 1 lần khi app start
func NewS3Provider(
	bucketName, region, apiKey, secret, endpoint, publicDomain string,
) *s3Provider {

	provider := &s3Provider{
		bucketName:   bucketName,
		region:       region,
		apiKey:       apiKey,
		secret:       secret,
		endpoint:     endpoint,
		publicDomain: publicDomain,
	}
	// Tạo AWS session
	// Session này sẽ được reuse cho tất cả các lần upload
	s3Session, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
		// Endpoint custom → dùng cho MinIO hoặc S3 self-hosted
		Endpoint:         aws.String(endpoint),
		S3ForcePathStyle: aws.Bool(true), // buộc sử dụng path style để tương thích với MinIO
		DisableSSL:       aws.Bool(true), // tắt SSL nếu sử dụng MinIO không có SSL
		Credentials: credentials.NewStaticCredentials(
			apiKey,
			secret,
			"",
		),
	})
	if err != nil {
		log.Fatalln(err)
	}

	provider.session = s3Session
	return provider
}

// SaveFileUploaded upload file lên S3 và trả về thông tin Image
//
// Params:
// - ctx: context request
// - data: nội dung file dạng []byte
// - dst: đường dẫn lưu trên S3 (vd: product/123.png)
//
// Return:
// - *common.Image: chứa URL public
// - error nếu upload thất bại
func (p *s3Provider) SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error) {
	// Tạo reader từ []byte
	fileBytes := bytes.NewReader(data)

	// Tự động detect mime-type (image/png, image/jpeg…)
	fileType := http.DetectContentType(data)

	// Thực hiện upload file lên S3 / MinIO
	_, err := s3.New(p.session).PutObject(&s3.PutObjectInput{
		Bucket: aws.String(p.bucketName),
		Key:    aws.String(dst), // Tên của file
		// ACL:         aws.String("private"),
		ContentType: aws.String(fileType), // Mime-type của file
		Body:        fileBytes,            // Nội dung file
	})
	if err != nil {
		return nil, err
	}

	// Tạo object Image để trả về cho business
	// URL public được build thủ công
	img := &common.Image{
		Url: fmt.Sprintf("%s/%s/%s",
			p.publicDomain, // vd: http://localhost:9000
			p.bucketName,   // vd: test
			dst,            // vd: product/abc.png
		),
		CloudName: "minio",
	}

	return img, nil
}
