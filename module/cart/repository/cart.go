package cartrepository

import "context"

type CartCleaner interface {
	DeleteCart(
		ctx context.Context,
		userId int,
	) error
}
