package model

import "food-delivery/common"

type User struct {
	common.SQLModel `json:",inline"`
	Email           string `json:"email" gorm:"column:email;"`
	Password        string `json:"-" gorm:"column:password;"`
	FBID            string `json:"fb_id" gorm:"column:fb_id;"`
	GGID            string `json:"gg_id" gorm:"column:gg_id;"`
	Salt            string `json:"-" gorm:"column:salt;"`
	LastName        string `json:"last_name" gorm:"column:last_name;"`
	FirstName       string `json:"first_name" gorm:"column:first_name;"`
	Phone           string `json:"phone" gorm:"column:phone;"`
	Role            string `json:"role" gorm:"column:role;"`
	Avatar          string `json:"avatar" gorm:"column:role;"`
}

func (User) TableName() string {
	return "users"
}

type UserCreate struct {
	common.SQLModel `json:",inline"`
	Email           string `json:"email" gorm:"column:email;"`
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

type UserUpdate struct {
	common.SQLModel `json:",inline"`
	Email           string `json:"email" gorm:"column:email;"`
	Password        string `json:"-" gorm:"column:password;"`
	FBID            string `json:"fb_id" gorm:"column:fb_id;"`
	GGID            string `json:"gg_id" gorm:"column:gg_id;"`
	Salt            string `json:"-" gorm:"column:salt;"`
	LastName        string `json:"last_name" gorm:"column:last_name;"`
	FirstName       string `json:"first_name" gorm:"column:first_name;"`
	Phone           string `json:"phone" gorm:"column:phone;"`
	Role            string `json:"role" gorm:"column:role;"`
	Avatar          string `json:"avatar" gorm:"column:role;"`
}
