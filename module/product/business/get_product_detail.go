package productbusiness

import (
	"OpenMarket/common"
	productmodel "OpenMarket/module/product/model"
	"context"
	"sync"
)

type getProductDetailBusiness struct {
	finder      GetProduct
	listvariant ListVariantRepo
	imagesRepo  GetImagesRepo
}

func NewGetProductDetailBusiness(
	finder GetProduct,
	listvariant ListVariantRepo,
	imagesRepo GetImagesRepo,
) *getProductDetailBusiness {
	return &getProductDetailBusiness{
		finder:      finder,
		listvariant: listvariant,
		imagesRepo:  imagesRepo,
	}
}

func (biz *getProductDetailBusiness) GetProductDetail(
	ctx context.Context,
	productID int,
) (*productmodel.ProductDetail, error) {
	p, err := biz.finder.GetProductById(ctx, productID)
	if err != nil || p == nil {
		return nil, common.ErrEntityNotFound(productmodel.EntityName, err)
	}

	var (
		variants []productmodel.VariantDetail
		urls     []string
		vErr     error
		iErr     error
		wg       sync.WaitGroup
	)

	wg.Add(2)

	go func() {
		defer wg.Done()
		variants, vErr = biz.listvariant.ListVariant(ctx, productID)
	}()

	go func() {
		defer wg.Done()
		urls, iErr = biz.imagesRepo.GetImages(ctx, productID)
	}()

	wg.Wait()

	if vErr != nil {
		return nil, common.ErrCannotListEntity("Variant", vErr)
	}
	if iErr != nil {
		return nil, common.ErrCannotListEntity("Images", iErr)
	}

	detail := &productmodel.ProductDetail{
		ID:          p.Id,
		Name:        p.Name,
		Description: p.Description,
		BasePrice:   p.BasePrice,
		Variants:    variants,
		Images:      urls,
	}

	return detail, nil
}
