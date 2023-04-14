package models


type SosialMedias struct {
	GormModel
	Name string `gorm:"not null" json:"name" form:"name" validate:"required-Your full name is required"`
	SosialMedia_Url string `gorm:"not null" json:"sosialMed" from:"sosialMed"`
	UserID uint `gorm:"primaryKey" json:"userid"`
	User *User


}