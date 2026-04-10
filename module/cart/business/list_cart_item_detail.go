package cartbusiness

import (
	"OpenMarket/common"
	cartmodel "OpenMarket/module/cart/model"
	"context"
)

type ListCartItemDetailRepo interface {
	FindCart(
		ctx context.Context,
		userId int,
	) (*cartmodel.Cart, error)

	ListCartItemsDetail(
		ctx context.Context,
		cartId int,
	) ([]cartmodel.CartItemDetailRow, error)
}

type listCartItemDetailBusiness struct {
	repo ListCartItemDetailRepo
}

func NewListCartItemDetailBusiness(repo ListCartItemDetailRepo) *listCartItemDetailBusiness {
	return &listCartItemDetailBusiness{repo: repo}
}

func (biz *listCartItemDetailBusiness) ListCartItemDetail(
	ctx context.Context,
	userId int,
) ([]cartmodel.CartItemDetail, error) {
	cart, err := biz.repo.FindCart(ctx, userId)
	if err != nil {
		if common.IsRecordNotFound(err) {
			return []cartmodel.CartItemDetail{}, nil
		}
		return nil, err
	}

	rows, err := biz.repo.ListCartItemsDetail(ctx, cart.Id)
	if err != nil {
		return nil, err
	}

	seen := make(map[int]*cartmodel.CartItemDetail)
	ordered := make([]*cartmodel.CartItemDetail, 0)

	for _, row := range rows {
		item, ok := seen[row.VariantId]
		if !ok {
			img := ""
			if row.ImageURL != nil {
				img = *row.ImageURL
			}
			item = &cartmodel.CartItemDetail{
				VariantId:     row.VariantId,
				Quantity:      row.Quantity,
				Price:         row.Price,
				StockQuantity: row.StockQuantity,
				Product: cartmodel.CartItemDetailProduct{
					Id:    row.ProductId,
					Name:  row.ProductName,
					Image: img,
				},
				Attributes: map[string]string{},
			}
			item.Product.Mask()
			seen[row.VariantId] = item
			ordered = append(ordered, item)
		}

		if row.AttributeName != nil && row.AttributeValue != nil {
			item.Attributes[*row.AttributeName] = *row.AttributeValue
		}
	}

	result := make([]cartmodel.CartItemDetail, 0, len(ordered))
	for _, it := range ordered {
		result = append(result, *it)
	}

	return result, nil
}
