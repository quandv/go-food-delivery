package restaurantstorage

import (
	"context"
	common "go-food-delivery/common"
	restaurantmodel "go-food-delivery/module/restaurant/model"
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

	/* Has 2 ways to handle "Gorm Error".
	* Refer doc: https://gorm.io/docs/error_handling.html
	 */
	if err := s.db.Where("id = ?", id).First(&data).Error; err != nil {
		return nil, common.ErrDB(restaurantmodel.EntityName, err)
	}

	// if result := s.db.Where("id = ?", id).First(&data); result.Error != nil {
	// 	return nil, common.ErrDB(restaurantmodel.EntityName, result.Error)
	// }

	log.Println("data in store: ", data)

	return &data, nil
}

func (s *sqlStore) Find(
	ctx context.Context,
	filter restaurantmodel.Filter,
	paging *common.Pagination,
	moreKeys ...string,
) ([]restaurantmodel.Restaurant, error) { // không cần con trỏ cho "[]restaurantmodel.Restaurant" bởi vì nó là slice, mà slice bản chất là con trỏ rồi nên nó có thể nhận giá trị "nil"
	db := s.db.Table(restaurantmodel.Restaurant{}.TableName())

	if f := filter; f.Name != "" {
		db.Where("name = ?", f.Name)
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(restaurantmodel.EntityName, err)
	}

	offset := (paging.Page - 1) * paging.Limit
	var data []restaurantmodel.Restaurant

	if err := db.Offset(offset).Limit(paging.Limit).Find(&data).Error; err != nil {
		return nil, common.ErrDB(restaurantmodel.EntityName, err)
	}

	return data, nil
}
