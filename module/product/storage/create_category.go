package productstorage

import "context"

func (s *sqlStore) CreateCategory(
	ctx context.Context,
	data map[string]interface{},
) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
