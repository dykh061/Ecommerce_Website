package productbusiness

import (
	"OpenMarket/common"
	productmodel "OpenMarket/module/product/model"
	"context"
	"errors"
	"sync"
)

type GetGalleriesRepo interface {
	GetGalleries(
		ctx context.Context,
		productID int,
	) ([]productmodel.GalleryItem, error)
}

type GetProductCategoryRepo interface {
	GetProductCategoryID(
		ctx context.Context,
		productID int,
	) (*int, error)
}

type getSellerProductDetailBusiness struct {
	sfinder      FindSellerByID
	pfinder      FindProductWithIDAndSellerID
	galleriesRepo GetGalleriesRepo
	categoryRepo GetProductCategoryRepo
}

func NewGetSellerProductDetailBusiness(
	sfinder FindSellerByID,
	pfinder FindProductWithIDAndSellerID,
	galleriesRepo GetGalleriesRepo,
	categoryRepo GetProductCategoryRepo,
) *getSellerProductDetailBusiness {
	return &getSellerProductDetailBusiness{
		sfinder:       sfinder,
		pfinder:       pfinder,
		galleriesRepo: galleriesRepo,
		categoryRepo:  categoryRepo,
	}
}

func (biz *getSellerProductDetailBusiness) GetSellerProductDetail(
	ctx context.Context,
	userID int,
	productID int,
) (*productmodel.SellerProductDetail, error) {
	// 1. Find seller by userID
	seller, err := biz.sfinder.FindActiveSellerWithUserID(ctx, userID)
	if err != nil {
		return nil, common.ErrForbidden(errors.New("user is not a seller"))
	}

	// 2. Find product with seller ownership check
	product, err := biz.pfinder.FindProductByIdWithSellerID(ctx, productID, seller.Id)
	if err != nil {
		return nil, common.ErrForbidden(errors.New("product does not belong to this seller"))
	}
	if product == nil {
		return nil, common.ErrEntityNotFound(productmodel.EntityName, errors.New("product not found"))
	}

	// 3. Get galleries and category in parallel
	var (
		galleries  []productmodel.GalleryItem
		categoryID *int
		gErr       error
		cErr       error
		wg         sync.WaitGroup
	)

	wg.Add(2)

	go func() {
		defer wg.Done()
		galleries, gErr = biz.galleriesRepo.GetGalleries(ctx, productID)
	}()

	go func() {
		defer wg.Done()
		categoryID, cErr = biz.categoryRepo.GetProductCategoryID(ctx, productID)
	}()

	wg.Wait()

	if gErr != nil {
		return nil, common.ErrCannotListEntity("Gallery", gErr)
	}
	if cErr != nil {
		return nil, common.ErrCannotListEntity("Category", cErr)
	}

	// 4. Build response
	productUID := common.NewUID(uint32(product.Id), common.DbTypeProduct, 1)
	sellerUID := common.NewUID(uint32(seller.Id), common.DbTypeSeller, 1)

	result := &productmodel.SellerProductDetail{
		ID:          productUID.String(),
		Name:        product.Name,
		Description: product.Description,
		BasePrice:   product.BasePrice,
		Status:      product.Status,
		CategoryID:  categoryID,
		SellerID:    sellerUID.String(),
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
		Images:      galleries,
	}

	return result, nil
}
