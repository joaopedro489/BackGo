package Models

import "github.com/jinzhu/gorm"

type User struct{
   gorm.Model
   Name string `json:"name"`
   Age int `json:"age"`
   Email string `json:"email" gorm:"unique"`
   Password string `json:"-" gorm:"not_null"`
}
