package restaurantstorage

func (s *store) ListDataWithCondition(
	context context.Context,
	filter *restaurantmodel.Filter,
)