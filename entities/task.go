package entities

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	TgId int
	Id   int
	Text string
}
