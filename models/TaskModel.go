package models

import (
	"MiddleTestTask/config"
	"fmt"
	"gorm.io/gorm"
)

type Tasks struct {
	gorm.Model
	Title       string `json:"title,omitempty" gorm:"type:varchar(255); not null"`
	Description string `json:"description,omitempty" gorm:"type:varchar(255)"`
}

func init() {
	db, err := config.Connection()
	if err != nil {
		fmt.Println(err)
	}
	if err := db.AutoMigrate(&Tasks{}); err != nil {
		fmt.Println(err)
	}
}
