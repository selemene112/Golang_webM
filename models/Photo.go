package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct{
	GormModel
	Title string `json:"title" form:"title" validate:"required-Title of your product is required"`
	Caption string `json:"caption" from:"caption" validate:"required-Description is required"`
	Photo_Url string `json:"photourl" from:"photourl"`
	UserID uint `gorm:"primaryKey" json:"userid"`
	User *User
	

}

func (p *Photo) BeforeCreate(tx *gorm.DB)(err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate 
		return 
	}

	err = nil
	return
}

func (p *Photo) BeforeUpdate(tx *gorm.DB)(err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate 
		return 
	}

	err = nil
	return
}





