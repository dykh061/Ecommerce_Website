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

type s3Provider struct {
	bucketName   string
	region       string
	apiKey       string
	secret       string
	endpoint     string
	publicDomain string
	session      *session.Session
}

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

	s3Session, err := session.NewSession(&aws.Config{
		Region:           aws.String(region),
		Endpoint:         aws.String(endpoint),
		S3ForcePathStyle: aws.Bool(true),
		DisableSSL:       aws.Bool(true),
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

func (p *s3Provider) SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error) {
	fileBytes := bytes.NewReader(data)
	fileType := http.DetectContentType(data)

	_, err := s3.New(p.session).PutObject(&s3.PutObjectInput{
		Bucket: aws.String(p.bucketName),
		Key:    aws.String(dst), // Tên của file
		// ACL:         aws.String("private"),
		ContentType: aws.String(fileType),
		Body:        fileBytes,
	})
	if err != nil {
		return nil, err
	}

	img := &common.Image{
		Url: fmt.Sprintf("%s/%s/%s",
			p.publicDomain,
			p.bucketName,
			dst,
		),
		CloudName: "minio",
	}

	return img, nil
}
