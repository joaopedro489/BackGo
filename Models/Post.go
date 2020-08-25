package Models
import (
  "github.com/jinzhu/gorm"
  DB "github.com/joaopedro489/BackGo/Database"
)
type Post struct{
  ID uint `json:"id" gorm:"primary_key"`
  Tilte string `json:"title"`
  Text string `json:text`
  PostId unit `json:"post_id"`
}

func (post *Post) createPost(postInput){
    post := md.Post{Name: postInput.Name, Age: postInput.Age, Email: postInput.Email, Password: postInput.Password}
    DB.DB.Create(&post)
}
func (post *Post) updatePost(postInput){
    DB.DB.Model(&post).Updates(postInput)

}
func (post *Post) deletePost(){
    DB.DB.Delete(&post)
}
