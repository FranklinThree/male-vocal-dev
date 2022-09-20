package structs

import (
	"gorm.io/gorm"
)

type hotVar struct {
	gorm.Model
	Name  string
	Value int
}

func (h hotVar) TableName() string {
	return "hotVar"
}
