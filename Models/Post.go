package Models

import (
  "github.com/jinzhu/gorm"
  "github.com/gin-gonic/gin"
  "net/http"
)
//atributos que irão para o BD com softDelete
type Post struct{
  gorm.Model
  Title string `json:"title"`
  Text string `json:"text"`
  User_id uint `json:"user_id"`
}

func (post *Post) GetPost(c *gin.Context){
  //apenas com o if o post será pego no DB
  if err := DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil{
    if err := DB.Where("id = ?", c.Param("post_id")).First(&post).Error; err != nil{
      c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
      return
    }
  }
}
func (post Post) CreatePost(c *gin.Context, user_id uint) Post{
  post = Post{Title: c.PostForm("title"), Text: c.PostForm("text"), User_id: user_id}
  DB.Create(&post)
  DB.Save(&post)
  return post
}
func (post *Post) UpdatePost(input interface{}){
  DB.Model(&post).Updates(input)
}
//usando softDelete
func (post *Post) DeletePost(){
  DB.Delete(&post)
}
