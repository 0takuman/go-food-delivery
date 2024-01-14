package restaurantstorage

import (
	"context"
	"food-delivery/common"

	restaurantmodel "food-delivery/modules/restaurant/model"
)

func (s *store) ListDataWithCondition(
	context context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]restaurantmodel.Restaurant, error) {
	var results []restaurantmodel.Restaurant
	db := s.db.Table(restaurantmodel.Restaurant{}.TableName())

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := db.Offset(paging.Page * paging.Limit).
		Limit(paging.Limit).Order("id desc").
		Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil

}
