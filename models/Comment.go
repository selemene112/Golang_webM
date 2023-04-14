package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	GormModel
	// Name string `gorm:"not null" json:"name" from:"name"`
	UserID uint `gorm:"primaryKey" json:"userid"`
	PhotoID uint `gorm:"primaryKey" json:"photoid"`
	User *User
	Message string    `gorm:"not null" json:"message" from:"message"`
	
}


func (p *Comment) BeforeCreate(tx *gorm.DB)(err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate 
		return 
	}

	err = nil
	return
}

func (p *Comment) BeforeUpdate(tx *gorm.DB)(err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate 
		return 
	}

	err = nil
	return
}

