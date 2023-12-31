package notestorage

import (
	"context"
	"food-delivery/common"
	nodemodel "food-delivery/modules/note/model"
)

func (s *store) ListDataWithCondition(context context.Context,
	filter *nodemodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]nodemodel.Notes, error) {
	var results []nodemodel.Notes
	db := s.db.Table(nodemodel.Notes{}.TableName())

	if f := filter; f != nil {
		if f.Status {
			db.Where("status=?", 1)
		} else {
			db.Where("status=?", 0)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := db.Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).Order("id desc").
		Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}
