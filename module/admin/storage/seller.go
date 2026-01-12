package adminstorage

import (
	"OpenMarket/common"
	adminmodel "OpenMarket/module/admin/model"
	"context"
	"strings"
)

// ListSellers returns paginated list of sellers for admin panel
func (s *sqlStore) ListSellers(
	ctx context.Context,
	filter *adminmodel.SellerFilter,
	paging *common.Paging,
) ([]adminmodel.SellerAdminView, error) {
	var sellers []adminmodel.SellerAdminView

	db := s.db.WithContext(ctx).
		Table("sellers s").
		Select(`
			s.id,
			s.user_id,
			s.shop_name,
			s.shop_description,
			s.shop_phone,
			s.status,
			s.created_at,
			s.updated_at,
			u.name as user_name,
			u.email as user_email,
			u.phone as user_phone
		`).
		Joins("LEFT JOIN users u ON u.id = s.user_id")

	// Apply filters
	if filter != nil {
		if filter.Status != nil {
			db = db.Where("s.status = ?", *filter.Status)
		}
		if filter.Keyword != nil && *filter.Keyword != "" {
			keyword := "%" + strings.ToLower(*filter.Keyword) + "%"
			db = db.Where("LOWER(s.shop_name) LIKE ? OR LOWER(u.email) LIKE ?", keyword, keyword)
		}
	}

	// Count total
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, err
	}
	paging.Total = total

	// Apply sorting
	orderBy := "s.created_at DESC"
	if filter != nil && filter.SortBy != "" {
		direction := "ASC"
		if filter.SortDesc {
			direction = "DESC"
		}
		orderBy = "s." + filter.SortBy + " " + direction
	}

	// Apply pagination
	offset := (paging.Page - 1) * paging.Limit
	if err := db.Order(orderBy).
		Offset(offset).
		Limit(paging.Limit).
		Scan(&sellers).Error; err != nil {
		return nil, err
	}

	return sellers, nil
}

// FindSellerById returns a seller by id for admin panel
func (s *sqlStore) FindSellerById(
	ctx context.Context,
	sellerId int,
) (*adminmodel.SellerAdminView, error) {
	var seller adminmodel.SellerAdminView

	if err := s.db.WithContext(ctx).
		Table("sellers s").
		Select(`
			s.id,
			s.user_id,
			s.shop_name,
			s.shop_description,
			s.shop_phone,
			s.status,
			s.created_at,
			s.updated_at,
			u.name as user_name,
			u.email as user_email,
			u.phone as user_phone
		`).
		Joins("LEFT JOIN users u ON u.id = s.user_id").
		Where("s.id = ?", sellerId).
		First(&seller).Error; err != nil {
		return nil, err
	}

	return &seller, nil
}

// UpdateSellerStatus updates seller status (lock/unlock)
func (s *sqlStore) UpdateSellerStatus(
	ctx context.Context,
	sellerId int,
	status int,
) error {
	return s.db.WithContext(ctx).
		Table("sellers").
		Where("id = ?", sellerId).
		Update("status", status).Error
}
