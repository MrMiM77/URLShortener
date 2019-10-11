package models

import (
	"github.com/jinzhu/gorm"
)

type URLMap struct {
	gorm.Model
	AbsoluteURL string
	ShortURL string
}