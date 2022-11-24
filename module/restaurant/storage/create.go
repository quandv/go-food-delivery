package restaurantstorage

import (
	"context"
	restaurantmodel "food-delivery/module/restaurant/model"
	"log"
)

func (s *sqlStore) CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := s.db.Create(&data).Error; err != nil {
		log.Panicln("s.db.Create => ", err)
		return err
	}
	return nil
}
