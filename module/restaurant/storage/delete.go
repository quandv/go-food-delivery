package restaurantstorage

import (
	"context"
	common "go-food-delivery/common"
	restaurantmodel "go-food-delivery/module/restaurant/model"
)

func (s *sqlStore) DeleteOne(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) error {
	var data restaurantmodel.Restaurant
	if err := s.db.Where(condition).Delete(&data).Error; err != nil {
		return common.ErrDB(restaurantmodel.EntityName, err)
	}
	return nil
}
