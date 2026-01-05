package productrepository

import (
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

func (repo *createVariantRepo) CreateVariantWithAttributes(
	ctx context.Context,
	productID int,
	data *productmodel.VariantCreate,
) error {
	data.ProductId = productID

	// mở transaction để tạo variant và các attribute values liên quan
	return repo.storage.WithTransaction(ctx, func(tx productstorage.TxStore) error {

		// 1 thực hiện tìm kiếm variant đã tồn tại với các attribute values này chưa
		existedVariant, err := tx.FindVariantWithAtributesValue(ctx, productID, data.AttributeValueIDs)
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}
			existedVariant = nil
		}

		// 2. nếu đã tồn tại thì chỉ cần điều chỉnh lại stock quantity
		if existedVariant != nil && existedVariant.Id > 0 {
			return tx.AdjustVariantStock(ctx, existedVariant.Id, data.StockQuantity)
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
}
