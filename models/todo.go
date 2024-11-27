package models

import "github.com/jinzhu/gorm"

type Todo struct {
	gorm.Model
	Task            string `gorm:"size:56;not null" form:"task" json:"task"`
	TaskDescription string `gorm:"size:255;not null" form:"task_description" json:"task_description"`
}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&Todo{})
}
