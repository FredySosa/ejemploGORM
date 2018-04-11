package models

import (
	"github.com/jinzhu/gorm"
)

type Actividad struct {
	gorm.Model
	idact int64 `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
}