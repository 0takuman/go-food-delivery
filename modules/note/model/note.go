package notemodel

type Notes struct {
	Id     int    `json:"id" gorm:"column:id"`
	Title  string `json:"title" gorm:"column:title"`
	Status bool   `json:"status" gorm:"column:status"`
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