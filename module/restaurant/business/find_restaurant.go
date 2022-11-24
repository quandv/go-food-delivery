package restaurantbusiness

import (
	restaurantmodel "food-delivery/module/restaurant/model"

	"golang.org/x/net/context"
)

type FindRestaurantStore interface {
	FindOne(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
	Find(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*[]restaurantmodel.Restaurant, error)
}

type findRestaurantBusiness struct {
	store FindRestaurantStore
}

func NewFindRestaurantBusiness(store FindRestaurantStore) *findRestaurantBusiness {
	return &findRestaurantBusiness{store: store}
}

func (business *findRestaurantBusiness) FindOne(
	context context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*restaurantmodel.Restaurant, error) {
	data, err := business.store.FindOne(context, condition)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (business *findRestaurantBusiness) Find(
	context context.Context,
	condition map[string]interface{},
	pagination map[string]interface{},
	// moreKeys ...string,
) (*[]restaurantmodel.Restaurant, error) {
	data, err := business.store.Find(context, condition)
	if err != nil {
		return nil, err
	}
	return data, nil
}
