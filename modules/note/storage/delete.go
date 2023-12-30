package notestorage

import (
	"context"
	nodemodel "food-delivery/modules/note/model"
)

func (s *store) Delete(context context.Context,
	id int,
) error {
	if err := s.db.Table(nodemodel.Notes{}.TableName()).Where("id=?", id).
		Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return err
	}
	return nil
}
