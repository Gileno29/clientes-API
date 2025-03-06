package models

import (
	"time"

	"gorm.io/gorm"
)

type Cliente struct {
	//gorm.Model
	Documento   string `gorm:"primaryKey;type:varchar(14)"`
	RazaoSocial string `gorm:"not null"`
	Blocklist   bool   `gorm:"default:false"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
