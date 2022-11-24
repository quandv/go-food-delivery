package restaurantstorage

import (
	"context"
	restaurantmodel "food-delivery/module/restaurant/model"
	"log"
)

func (s *sqlStore) FindOne(
	ctx context.Context,
	condition map[string]interface{}, // map có key là string và value là bất kỳ
	moreKeys ...string,
) (*restaurantmodel.Restaurant, error) {
	// why should pointer at here?
	// Trong trường hợp bị lỗi, chúng ta sẽ "return nil, error" và giá trị nil sẽ bị lỗi nếu định nghĩa type trả về ko phải là con trỏ
	// Trong trường hợp ko phải là con trỏ, chúng ta sẽ phải return restaurantmodel.Restaurant{} = một struct rỗng => zeroed value (và vẫn mất bộ nhớ)
	// => giúp giảm thiểu bộ nhớ

	var data restaurantmodel.Restaurant
	id := condition["id"]

	if err := s.db.Where("id = ?", id).First(&data); err.Error != nil {
		log.Println("err in store: ", err)
		return nil, err.Error
	}

	log.Println("data in store: ", data)

	return &data, nil
}

func (s *sqlStore) Find(
	ctx context.Context,
	condition map[string]interface{},
	// pagination map[string]interface{},
	// moreKeys ...string,
) (*[]restaurantmodel.Restaurant, error) {
	db := s.db.Table(restaurantmodel.Restaurant{}.TableName()).Where(condition)

	// page := pagination["page"]
	// limit := pagination["limit"]

	// log.Println(page)
	// log.Println(limit)
	// offset := (page - 1) * limit
	var data []restaurantmodel.Restaurant

	if err := db.Find(&data); err.Error != nil {
		return nil, err.Error
	}

	return &data, nil
}
