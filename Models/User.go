package Models
import (
  "github.com/jinzhu/gorm"
  DB "github.com/joaopedro489/BackGo/Database"
)

type User struct{
   ID uint `json:"id" gorm:"primary_key"`
   Name string `json:"name"`
   Age uint `json:"age"`
   Email string `json:"email" gorm:"unique"`
   Password string `json::"password" gorm:"not_null"`
   Posts []Post `json:"posts" gorm:"foreignkey:UserId"`
}

func (user *User) createUser(userInput){
    user := md.User{Name: userInput.Name, Age: userInput.Age, Email: userInput.Email, Password: userInput.Password}
    DB.DB.Create(&user)
}
func (user *User) updateUser(userInput){
    DB.DB.Model(&user).Updates(userInput)

}
func (user *User) deleteUser(){
    DB.DB.Delete(&user)
}
