package notemodel

import "food-delivery/common"

type Notes struct {
	common.SQLModel `json:",inline"`
	Title           string `json:"title" gorm:"column:title"`
}

func (Notes) TableName() string {
	return "notes"
}

type NotesUpdate struct {
	Id     int     `json:"id" gorm:"column:id"`
	Title  *string `json:"title" gorm:"column:title"`
	Status bool    `json:"status" gorm:"column:status"`
}

func (NotesUpdate) TableName() string {
	return Notes{}.TableName()
}
