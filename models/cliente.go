package models

import "gorm.io/gorm"

type Cliente struct {
	gorm.Model
	Documento string `gorm:"unique;not null"` 
	Nome      string `gorm:"not null"`       
	Blocklist bool   `gorm:"default:false"`   
}