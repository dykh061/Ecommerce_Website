package userstorage

import "context"

func (s *sqlStore) DeleteAddress(
	ctx context.Context,
	id, userId int,
) error {
	if err := s.db.Table("user_addresses").
		Where("id = ? AND user_id = ?", id, userId).
		Delete(nil).Error; err != nil {
		return err
	}
	return nil
}
