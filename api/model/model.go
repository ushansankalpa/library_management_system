package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Model struct {
	ID        int64      `gorm:"primary_key" json:"id" swaggerignore:"true"`
	CreatedAt time.Time  `swaggerignore:"true"`
	UpdatedAt time.Time  `swaggerignore:"true"`
	DeletedAt *time.Time `sql:"index" swaggerignore:"true"`
}

func InitModels(db *gorm.DB) {
	db.AutoMigrate(&Members{})
	db.AutoMigrate(&Books{})
	db.AutoMigrate(&Publishers{})

}
