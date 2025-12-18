package productstorage

import (
	"OpenMarket/common"
	productmodel "OpenMarket/module/product/model"
	"context"
)

func (s *sqlStore) ListDataWithCondition(
	context context.Context,
	filter *productmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]productmodel.Product, error) {
	var result []productmodel.Product

	db := s.db.Table(productmodel.Product{}.TableName())

	if f := filter; f != nil {
		if f.SellerID != nil {
			db = db.Where("seller_id = ?", *f.SellerID)
		}
		if f.Status != nil {
			db = db.Where("status = ?", *f.Status)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	offset := (paging.Page - 1) * paging.Limit

	if err := db.Offset(offset).
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).
		Error; err != nil {
		return nil, err
	}
	return result, nil
}
