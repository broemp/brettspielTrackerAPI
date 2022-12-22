package entity

import (
	"gorm.io/gorm"
)

type Collection struct {
	gorm.Model
	Username string      `json:"username" gorm:"type:varchar(64)"`
	Games    []Boardgame `json:"boardgames" gorm:"many2many:collection_boardgames;"`
}
