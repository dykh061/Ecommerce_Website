package orderstorage

import (
	"OpenMarket/common"
	ordermodel "OpenMarket/module/order/model"
	"context"
)

func (s *sqlStore) ListOrders(
	ctx context.Context,
	filter *ordermodel.FilterOrder,
	paging *common.Paging,
) ([]ordermodel.Order, error) {
	var result []ordermodel.Order
	db := s.db.Table(ordermodel.Order{}.TableName())
	if f := filter; f != nil {
		if f.UserId != nil {
			db = db.Where("user_id = ?", *f.UserId)
		}
		if f.Status != nil {
			db = db.Where("LOWER(status) LIKE LOWER(?)", *f.Status)
		}
		if f.MaxPrice != nil {
			db = db.Where("total_amount <= ?", *f.MaxPrice)
		}
		if f.MinPrice != nil {
			db = db.Where("total_amount >= ?", *f.MinPrice)
		}
	}

	dbCount := db

	if err := dbCount.Count(&paging.Total).Error; err != nil {
		return nil, err
	}
	if err := db.Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&result).
		Error; err != nil {
		return nil, err
	}
	return result, nil
}
