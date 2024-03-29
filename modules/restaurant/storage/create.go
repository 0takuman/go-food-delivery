package restaurantstorage

import (
	"context"
	restaurantmodel "food-delivery/modules/restaurant/model"
)

func (s *store) CreateRestaurant(context context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
