package notestorage

import "context"

func (s *store) FindDataWithCondition(context context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*restaurantmodel.Restaurant, error) {
	var restaurant restaurantmodel.Restaurant
	if err := s.db.Where(condition).First(&restaurant).Error; err != nil {
		return nil, err
	}
	return &restaurant, nil
}
