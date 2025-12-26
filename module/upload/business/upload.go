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

type uploadBusiness struct {
	provider uploadprovider.UploadProvider
}

func NewUploadBusiness(provider uploadprovider.UploadProvider) *uploadBusiness {
	return &uploadBusiness{provider: provider}
}

func (b *uploadBusiness) UploadFile(ctx context.Context, data []byte, folder, filename string) (*common.Image, error) {
	fileBytes := bytes.NewBuffer(data)
	w, h, err := getImageDimension(fileBytes)
	if err != nil {
		return nil, err
	}
	fileExt := filepath.Ext(filename)
	filename = fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt)
	img, err := b.provider.SaveFileUploaded(ctx, data, fmt.Sprintf("%s/%s", folder, filename))
	if err != nil {
		return nil, err
	}

	img.Width = w
	img.Height = h
	img.Extension = fileExt
	return img, nil
}

func getImageDimension(reader io.Reader) (int, int, error) {
	img, _, err := image.DecodeConfig(reader)
	if err != nil {
		log.Println("Failed to decode image:", err)
		return 0, 0, err
	}
	return img.Width, img.Height, nil
}
