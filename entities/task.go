package entities

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	TgId int64 `gorm:"column:TgId"`
	Id   int
	Text string
}
