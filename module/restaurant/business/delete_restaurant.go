package restaurantbusiness

import (
	"context"
	"errors"
	restaurantmodel "food-delivery/module/restaurant/model"
)

type DeleteRestaurantStore interface {
	DeleteOne(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) error
	FindOne(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
}

type deleteRestaurantBusiness struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBusiness(store DeleteRestaurantStore) *deleteRestaurantBusiness {
	return &deleteRestaurantBusiness{store: store}
}

func (business *deleteRestaurantBusiness) DeleteOne(
	context context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) error {
	oldData, err := business.store.FindOne(context, condition)

	if err != nil {
		return err
	}

	if oldData != nil {
		return errors.New("restauran has been deleted")
	}

	if err := business.store.DeleteOne(context, condition); err != nil {
		return err
	}
	return nil
}
