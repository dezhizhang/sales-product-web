package model

import "time"

type BaseModel struct {
	Id        string    `gorm:"primaryKey;type:varchar(200)" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:update_at" json:"update_at"`
	DeletedAt time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	IsDeleted bool      `json:"isDeleted"`
}

type ResponseData struct {
	Code  int
	Msg   string
	Data  any
	Total int32
}
