package notestorage

import (
	"context"
	nodemodel "food-delivery/modules/note/model"
)

func (s *store) FindDataWithCondition(context context.Context, 
		condition map[string]interface{},
		moreKeys ...string,
	) (*nodemodel.Notes, error) {
	var note nodemodel.Notes
	if err := s.db.Where(condition).First(&note).Error; err != nil {
		return nil, err
	}
	return &note, nil
}