package restaurantbusiness

import (
	"context"
	restaurantmodel "food-delivery/module/restaurant/model"
	"log"
)

// Golang convention inteface: Inteface được khai báo ở nơi dùng nó
type CreateRestaurantStore interface {
	CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error
}

type createRestaurantBusiness struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantBusiness(store CreateRestaurantStore) *createRestaurantBusiness {
	return &createRestaurantBusiness{store: store}
}

func (business *createRestaurantBusiness) CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	// business logic code

	if err := business.store.CreateRestaurant(ctx, data); err != nil {
		log.Println("business.CreateRestaurant => ", err)
		return err
	}

	return nil
}
