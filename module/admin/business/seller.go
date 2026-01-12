package adminbusiness

import (
	"OpenMarket/common"
	adminmodel "OpenMarket/module/admin/model"
	"context"
	"errors"
)

type ListSellersStore interface {
	ListSellers(
		ctx context.Context,
		filter *adminmodel.SellerFilter,
		paging *common.Paging,
	) ([]adminmodel.SellerAdminView, error)
}

type listSellersBusiness struct {
	store ListSellersStore
}

func NewListSellersBusiness(store ListSellersStore) *listSellersBusiness {
	return &listSellersBusiness{store: store}
}

func (biz *listSellersBusiness) ListSellers(
	ctx context.Context,
	filter *adminmodel.SellerFilter,
	paging *common.Paging,
) ([]adminmodel.SellerAdminView, error) {
	sellers, err := biz.store.ListSellers(ctx, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity("Seller", err)
	}

	for i := range sellers {
		sellers[i].Mask()
	}

	return sellers, nil
}

// =========================================

type GetSellerStore interface {
	FindSellerById(ctx context.Context, sellerId int) (*adminmodel.SellerAdminView, error)
}

type getSellerBusiness struct {
	store GetSellerStore
}

func NewGetSellerBusiness(store GetSellerStore) *getSellerBusiness {
	return &getSellerBusiness{store: store}
}

func (biz *getSellerBusiness) GetSeller(
	ctx context.Context,
	sellerId int,
) (*adminmodel.SellerAdminView, error) {
	seller, err := biz.store.FindSellerById(ctx, sellerId)
	if err != nil {
		return nil, common.ErrEntityNotFound("Seller", err)
	}

	seller.Mask()
	return seller, nil
}

// =========================================

type UpdateSellerStatusStore interface {
	FindSellerById(ctx context.Context, sellerId int) (*adminmodel.SellerAdminView, error)
	UpdateSellerStatus(ctx context.Context, sellerId int, status int) error
}

type updateSellerStatusBusiness struct {
	store UpdateSellerStatusStore
}

func NewUpdateSellerStatusBusiness(store UpdateSellerStatusStore) *updateSellerStatusBusiness {
	return &updateSellerStatusBusiness{store: store}
}

func (biz *updateSellerStatusBusiness) UpdateStatus(
	ctx context.Context,
	sellerId int,
	data *adminmodel.SellerStatusUpdate,
) error {
	// Check if seller exists
	seller, err := biz.store.FindSellerById(ctx, sellerId)
	if err != nil {
		return common.ErrEntityNotFound("Seller", err)
	}

	// Validate status transition
	if seller.Status == data.Status {
		if data.Status == adminmodel.SellerStatusActive {
			return common.NewCustomError(errors.New("seller is already active"), "seller is already active", "ErrSellerAlreadyActive")
		}
		return common.NewCustomError(errors.New("seller is already inactive"), "seller is already locked", "ErrSellerAlreadyLocked")
	}

	// Update status
	if err := biz.store.UpdateSellerStatus(ctx, sellerId, data.Status); err != nil {
		return common.ErrCannotUpdateEntity("Seller", err)
	}

	return nil
}
