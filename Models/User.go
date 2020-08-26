package Models

import (
  "github.com/jinzhu/gorm"
  "github.com/gin-gonic/gin"
  "strconv"
  "net/http"
)
//atributos que irão para o BD com softDelete
type User struct{
   gorm.Model
   Name string `json:"name"`
   Age int `json:"age"`
   Email string `json:"email" gorm:"unique"`
   //json como "-" para não mostrar a senha
   Password string `json:"-" gorm:"not_null"`
}

func (user *User) GetUser(c *gin.Context){
  //apenas com o if o user será pego no DB
  if err := DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil{
    if err := DB.Where("id = ?", c.Param("user_id")).First(&user).Error; err != nil{
      c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
      return
    }
  }
}
func (user User) CreateUser(c *gin.Context) User{
  age,_ := strconv.Atoi(c.PostForm("age"));
  user = User{Name: c.PostForm("name"), Age: age, Email: c.PostForm("email"), Password: c.PostForm("password")}
  DB.Create(&user)
  DB.Save(&user)
  return user
}
func (user *User) UpdateUser(input interface{}){
  DB.Model(&user).Updates(input)
}
//usando HardDelete
func (user *User) DeleteUser(){
  DB.Unscoped().Delete(&user)
}
func (user *User) GetPosts(posts *[]Post){
  DB.Where("user_id=?", user.ID).Find(&posts)
}
