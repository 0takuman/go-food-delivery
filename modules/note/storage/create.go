package notestorage

import (
	"context"
	notemodel "food-delivery/modules/note/model"
)

func (s *store) CreateNote(context context.Context, data *notemodel.Notes) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
