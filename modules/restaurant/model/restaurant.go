package restaurantmodel

import "food-delivery/common"

const EntityName = "Restaurant"

type Restaurant struct {
	common.SQLModel  `json:",inline"`
	OwnerId          int            `json:"owner_id" gorm:"column:owner_id;"`
	Name             string         `json:"name" gorm:"column:name;"`
	Address          string         `json:"address" gorm:"column:addr;"`
	CityId           int            `json:"city_id" gorm:"column:city_id;"`
	Lat              float64        `json:"lat" gorm:"column:lat;"`
	Lng              float64        `json:"lng" gorm:"column:lng;"`
	Cover            *common.Images `json:"cover" gorm:"column:cover;"`
	Logo             *common.Image  `json:"logo" gorm:"column:logo;"`
	ShippingFeePerKm int            `json:"shipping_fee_per_km" gorm:"column:shipping_fee_per_km;"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

func (r *Restaurant) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeRestaurant)
}

type RestaurantCreate struct {
	common.SQLModel `json:",inline"`
	OwnerId         int            `json:"owner_id" gorm:"column:owner_id;"`
	Name            string         `json:"name" gorm:"column:name;"`
	Address         string         `json:"address" gorm:"column:addr;"`
	UserId          int            `json:"-" gorm:"column:user_id;"`
	Cover           *common.Images `json:"cover" gorm:"column:cover;"`
	Logo            *common.Image  `json:"logo" gorm:"column:logo;"`
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

func (r *RestaurantCreate) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeRestaurant)
}

type RestaurantUpdate struct {
	common.SQLModel `json:",inline"`
	OwnerId         int            `json:"owner_id" gorm:"column:owner_id;"`
	Name            string         `json:"name" gorm:"column:name;"`
	Cover           *common.Images `json:"cover" gorm:"column:cover;"`
	Logo            *common.Image  `json:"logo" gorm:"column:logo;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}
