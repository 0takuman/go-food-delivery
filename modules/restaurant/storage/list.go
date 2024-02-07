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

	// Fake Cursor

	if v := paging.FakeCursor; v != "" {
		uid, err := common.FromBase58(v)
		if err != nil {
			return nil, common.ErrDB(err)
		}

		db = db.Where("id < ?", uid.GetLocalID())
	} else {
		offset := (paging.Page - 1) * paging.Limit
		db = db.Offset(offset)
	}

	if err := db.
		Limit(paging.Limit).Order("id desc").
		Find(&results).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if len(results) > 0 {
		last := results[len(results)-1]
		last.Mask(false)
		paging.NextCursor = last.FakeId.String()
	}

	return results, nil

}
