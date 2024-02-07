package restaurantstorage

import (
	"context"
	"gorm.io/gorm"
)

type store struct {
	db *gorm.DB
}

// Delete implements restaurantbiz.DeleteRestaurantStore.
func (*store) Delete(ctx context.Context, id int) error {
	panic("unimplemented")
}

func NewSQLStore(db *gorm.DB) *store {
	return &store{db: db}
}
