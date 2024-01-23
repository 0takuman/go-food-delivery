package restaurantbiz

import (
	"context"
	"errors"
	"food-delivery/common"
	restaurantmodel "food-delivery/modules/restaurant/model"
)

type CreateRestaurantStore interface {
	CreateRestaurant(context context.Context, data *restaurantmodel.RestaurantCreate) error
}

type createRestaurantBiz struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantBiz(store CreateRestaurantStore) *createRestaurantBiz {
	return &createRestaurantBiz{store: store}
}

func (biz *createRestaurantBiz) CreateRestaurant(context context.Context, data *restaurantmodel.RestaurantCreate) error {
	if data.Name == "" {
		return errors.New("name is empty")
	}

	if err := biz.store.CreateRestaurant(context, data); err != nil {
		return common.ErrCannotCreateEntity(restaurantmodel.EntityName, err)
	}

	return nil

}
