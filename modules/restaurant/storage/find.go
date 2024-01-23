package restaurantstorage

import (
	"context"
	"food-delivery/common"
	restaurantmodel "food-delivery/modules/restaurant/model"

	"gorm.io/gorm"
)

func (s *store) FindDataWithCondition(context context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*restaurantmodel.Restaurant, error) {
	var restaurant restaurantmodel.Restaurant
	if err := s.db.Where(condition).First(&restaurant).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &restaurant, nil
}
