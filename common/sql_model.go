package common

import "time"


type SQLModel struct {
	Id int `json:"-" gorm: "column:id;"`
	// FakeId *UID `json:"id" gorm: "-"`
	Status   int        `json:"status" gorm: "column:status;default:1"`
	CreateAt *time.Time `json:"create_at;omitempty" gorm: "column:create_at"`
	UpdateAt *time.Time `json:"update_at;omitempty" gorm: "column:update_at"`
}
