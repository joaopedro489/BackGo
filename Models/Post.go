package Models

import "github.com/jinzhu/gorm"

type Post struct{
  gorm.Model
  Tilte string `json:"title"`
  Text string `json:"text"`
  User_id uint `json:"user_id"`
}
