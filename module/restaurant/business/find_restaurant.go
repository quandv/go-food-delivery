package restaurantbusiness

import (
	common "go-food-delivery/common"
	restaurantmodel "go-food-delivery/module/restaurant/model"

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
		filter restaurantmodel.Filter,
		paging *common.Pagination,
		moreKeys ...string,
	) ([]restaurantmodel.Restaurant, error)
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
	filter restaurantmodel.Filter,
	paging *common.Pagination,
	moreKeys ...string,
) ([]restaurantmodel.Restaurant, error) {
	data, err := business.store.Find(context, filter, paging)
	if err != nil {
		return nil, err
	}
	return data, nil
}
