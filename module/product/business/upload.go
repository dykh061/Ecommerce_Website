package productbusiness

import (
	"OpenMarket/common"
	"context"
)

type ImageUploader interface {
	UploadFile(
		ctx context.Context,
		data []byte,
		folder string,
		filename string,
	) (*common.Image, error)
}
