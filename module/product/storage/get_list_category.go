package productstorage

import (
	"OpenMarket/common"
	productmodel "OpenMarket/module/product/model"
	"context"
)

func (s *sqlStore) GetListCategory(
	ctx context.Context,
	filter *productmodel.FilterCategory,
	paging *common.Paging,
	moreKeys ...string,
) ([]productmodel.Category, error) {
	var result []productmodel.Category
	db := s.db.Table(productmodel.Category{}.TableName())
	if f := filter; f != nil {
		if f.Search != nil {
			db = db.Where("LOWER(name) LIKE LOWER(?)", "%"+*f.Search+"%")
		}
		if f.ParentID != nil {
			db = db.Where("parent_id = ?", *f.ParentID)
		}
	}
	dbCount := db

	if err := dbCount.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}
	if err := db.Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&result).
		Error; err != nil {
		return nil, err
	}
	return result, nil

}
