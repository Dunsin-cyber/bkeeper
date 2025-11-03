package models

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	FirstName *string `gorm:"type:varchar(200)"`
	LastName  *string `gorm:"type:varchar(200)"`
	Email     string  `gorm:"type:varchar(100);unique;not null"`
	Gender    *string `gorm:"varchar(10)"`
	Password  string  `gorm:"varchar(255);not null"`
}

func (u UserModel) TableName() string {
	return "users"
}
