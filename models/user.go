package models

import (
	"encoding/json"
	"final/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)



type User struct{
GormModel
Username string `gorm:"not null" json:"username" form:"username" validate:"required-Your full name is required"`//validate
Email string `gorm:"not null;uniqueIndex" json:"email" form:"email" validate:"required-Your email is required, email~invaild email format"`
Password string `gorm:"not null" json:"password" validate:"required,min=6"`//required-Your password is required, midstringlength(6)~password has to have a minimum length of 6 characters//form:"password"
Age string `gorm:"not null" json:"age" from:"age"`
Photos  []Photo `json:"photos"`
SosialMedia []SosialMedias `json:"sosialmedia"`




}

func (u *User) BeforeCreate(tx *gorm.DB)(err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate 
		return 
	}
	u.Password = helpers.HashPass(u.Password)
	err = nil
	

	userJson := `{}`
  var user User
  err = json.Unmarshal([]byte(userJson), &user)
  if err != nil {
    panic(err)
  }
  return
}

