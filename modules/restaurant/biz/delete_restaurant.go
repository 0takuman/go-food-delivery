package restaurantbiz

import (
	"context"
	"food-delivery/common"
	restaurantmodel "food-delivery/modules/restaurant/model"
)

type DeleteRestaurantStore interface {
	FindDataWithCondition(ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string) (*restaurantmodel.Restaurant, error)
	Delete(ctx context.Context, id int) error
}

type deleteRestaurantBiz struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store: store}
}

func (biz *deleteRestaurantBiz) DeleteRestaurant(ctx context.Context, id int) error {
	oldData, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrCannotDeleteEntity(restaurantmodel.EntityName, err)
	}

	if oldData.Status != 1 {
		return common.ErrEntityDeleted(restaurantmodel.EntityName, nil)
	}

	if err := biz.store.Delete(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(restaurantmodel.EntityName, err)
	}

	return nil
}
