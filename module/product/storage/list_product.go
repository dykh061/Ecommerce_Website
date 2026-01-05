package productstorage

import (
	"OpenMarket/common"
	productmodel "OpenMarket/module/product/model"

	"context"
)

func (s *sqlStore) ListProduct(
	ctx context.Context,
	filter *productmodel.Filter,
	paging *common.Paging,
) ([]productmodel.ProductListItem, error) {
	var result []productmodel.ProductListItem

	db := s.db.WithContext(ctx).
		Table(productmodel.Product{}.TableName() + " p").
		Select(`p.id,p.name,p.base_price,g.image_url`).
		Joins(`left join galleries g 
				on p.id = g.product_id and g.is_main  = true`)

	if f := filter; f != nil {
		if f.Status != nil {
			db = db.Where("p.status = ?", *f.Status)
		} else {
			db = db.Where("p.status = ?", common.SystemStatusActive)
		}
		if f.SellerID != nil {
			db = db.Where("p.seller_id = ?", *f.SellerID)
		}
		if f.CategoryID != nil {
			db = db.Joins("JOIN product_categories pc ON pc.product_id = p.id").
				Where("pc.category_id = ?", *f.CategoryID)
		}
		if f.Search != nil {
			db = db.Where("lower(p.name) like lower(?) or lower(p.description) like lower(?)",
				"%"+*f.Search+"%",
				"%"+*f.Search+"%")
		}
		if f.MinPrice != nil {
			db = db.Where("p.base_price >= ?", *f.MinPrice)
		}
		if f.MaxPrice != nil {
			db = db.Where("p.base_price <= ?", *f.MaxPrice)
		}
	}
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}
	if err := db.Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Scan(&result).
		Error; err != nil {
		return nil, err
	}
	return result, nil
}
