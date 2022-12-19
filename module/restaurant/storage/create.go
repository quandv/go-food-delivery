package restaurantstorage

import (
	"context"
	common "go-food-delivery/common"
	restaurantmodel "go-food-delivery/module/restaurant/model"
	"log"
)

func (s *sqlStore) CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := s.db.Create(&data).Error; err != nil {
		log.Panicln("s.db.Create => ", err)
		return common.ErrDB(restaurantmodel.EntityName, err)
	}
	return nil
}
