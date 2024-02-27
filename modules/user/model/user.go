package usermodel

import (
	"errors"
	"food-delivery/common"
)

const EntityName = "User"

type User struct {
	common.SQLModel `json:",inline"`
	Email           string        `json:"email" gorm:"column:email;"`
	Password        string        `json:"-" gorm:"column:password;"`
	FBID            string        `json:"fb_id" gorm:"column:fb_id;"`
	GGID            string        `json:"gg_id" gorm:"column:gg_id;"`
	Salt            string        `json:"-" gorm:"column:salt;"`
	LastName        string        `json:"last_name" gorm:"column:last_name;"`
	FirstName       string        `json:"first_name" gorm:"column:first_name;"`
	Phone           string        `json:"phone" gorm:"column:phone;"`
	Role            string        `json:"role" gorm:"column:role;"`
	Avatar          *common.Image `json:"avatar" gorm:"column:role;"`
}

func (User) TableName() string {
	return "users"
}

type UserCreate struct {
	common.SQLModel `json:",inline"`
	Email           string `json:"email" gorm:"column:email;"`
	Password        string `json:"-" gorm:"column:password;"`
	Salt            string `json:"-" gorm:"column:salt;"`
	Role            string `json:"role" gorm:"column:role;"`
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

func (u *UserCreate) Mask(isAdminOrOwner bool) {
	u.GenUID(common.DbTypeUser)
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

type UserLogin struct {
	Email    string `json:"email" form:"email" gorm:"column:email;"`
	Password string `json:"password" form:"password" gorm:"column:password;"`
}

func (UserLogin) TableName() string {
	return User{}.TableName()
}

var (
	ErrUsernameOrPasswordInvalid = common.NewCustomError(
		errors.New("username or password invalid"),
		"username or password invalid",
		"ErrUsernameOrPasswordInvalid",
	)

	ErrEmailExisted = common.NewCustomError(
		errors.New("email has already existed"),
		"Email has already existed",
		"ErrEmailExisted",
	)
)
