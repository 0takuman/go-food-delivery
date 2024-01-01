package notemodel

import "food-delivery/common"

type Notes struct {
	common.SQLModel `json:",inline"`
	Title           string `json:"title" gorm:"column:title"`
}

func (Notes) TableName() string {
	return "notes"
}

func (n *Notes) Mask(isAdminOrOwner bool) {
	n.GenUID(common.DbTypeNote)
}

type NotesUpdate struct {
	common.SQLModel `json:",inline"`
	Title           *string `json:"title" gorm:"column:title"`
}

func (NotesUpdate) TableName() string {
	return Notes{}.TableName()
}
