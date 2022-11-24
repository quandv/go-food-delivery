package restaurantstorage

import (
	"context"
	restaurantmodel "food-delivery/module/restaurant/model"
)

func (s *sqlStore) DeleteOne(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) error {
	var data restaurantmodel.Restaurant
	if err := s.db.Where(condition).Delete(&data); err != nil {
		return err.Error
	}
	return nil
}
