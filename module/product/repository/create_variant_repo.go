package productrepository

import (
	"OpenMarket/common"
	productmodel "OpenMarket/module/product/model"
	productstorage "OpenMarket/module/product/storage"
	"context"
	"errors"

	"gorm.io/gorm"
)

type CreateVariantStorage interface {
	WithTransaction(
		ctx context.Context,
		fn func(tx productstorage.TxStore) error,
	) error
}

type createVariantRepo struct {
	storage CreateVariantStorage
}

func NewCreateVariantRepo(storage CreateVariantStorage) *createVariantRepo {
	return &createVariantRepo{storage: storage}
}

// ErrDuplicateVariant is returned when a variant with same attributes already exists
var ErrDuplicateVariant = errors.New("variant with these attributes already exists")

func (repo *createVariantRepo) CreateVariantWithAttributes(
	ctx context.Context,
	productID int,
	data *productmodel.VariantCreate,
) (*productmodel.VariantCreate, error) {
	data.ProductId = productID

	// mở transaction để tạo variant và các attribute values liên quan
	err := repo.storage.WithTransaction(ctx, func(tx productstorage.TxStore) error {

		// 1 thực hiện tìm kiếm variant đã tồn tại với các attribute values này chưa (chỉ variant active)
		existedVariant, err := tx.FindVariantWithAtributesValue(ctx, productID, data.AttributeValueIDs)
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}
			existedVariant = nil
		}

		// 2. nếu đã tồn tại VÀ variant đó đang active → trả lỗi ErrVariantAlreadyExists
		if existedVariant != nil && existedVariant.Id > 0 && existedVariant.Status == common.SystemStatusActive {
			return ErrDuplicateVariant
		}

		// 3. nếu chưa tồn tại thì thực hiện tạo mới variant
		if err := tx.CreateVariant(ctx, data); err != nil {
			return err
		}

		// 4. tạo các bản ghi trong bảng variant_attribute_values
		rows := make([]productmodel.VariantAttributeValue, 0, len(data.AttributeValueIDs)) // chuẩn bị dữ liệu

		// gán variant_id và attribute_value_id
		for _, avID := range data.AttributeValueIDs {
			rows = append(rows, productmodel.VariantAttributeValue{
				VariantID:        data.Id,
				AttributeValueID: avID,
			})
		}

		return tx.CreateVariantAttributeValues(ctx, rows)
	})

	if err != nil {
		if errors.Is(err, ErrDuplicateVariant) {
			return nil, common.ErrVariantAlreadyExists(err)
		}
		return nil, err
	}

	data.Mask()
	return data, nil
}
