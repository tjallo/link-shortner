package models

import (
	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	OriginalLink string
	ShortLink    string `gorm:"index:idx_name,unique"`
}
