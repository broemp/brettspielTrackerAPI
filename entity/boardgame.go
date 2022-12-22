package entity

import (
	"gorm.io/gorm"
)

type Boardgame struct {
	gorm.Model
	BGGID       int     `json:"bggid"`
	Name        string  `json:"name" binding:"required,min=1,max=128" gorm:"type:varchar(128)"`
	Description string  `json:"description" binding:"max=10000" gorm:"type:varchar(10000)"`
	ImageUrl    string  `json:"imageUrl" binding:"max=512" gorm:"type:varchar(512)"`
	ReleaseYear string  `json:"releaseYear" gorm:"type:varchar(5)"`
	Rating      float64 `json:"rating" gorm:"type:decimal(10,1)"`
	MinPlayer   int     `json:"minplayer" gorm:"type:int"`
	MaxPlayer   int     `json:"maxplayer" gorm:"type:int"`
	MinPlaytime int     `json:"minplaytime" gorm:"type:int"`
	MaxPlaytime int     `json:"maxplaytime" gorm:"type:int"`
	Complexity  float64 `json:"complexity" gorm:"type:decimal(10,1)"`
}
