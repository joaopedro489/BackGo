package Models
import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)
var DB *gorm.DB

func ConnectDatabase(){
  db,err := gorm.Open("mysql", "root@/demo?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
  panic(err)
 }
 db.AutoMigrate(&User{}, &Post{})
 db.Model(&Post{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
 DB = db
}
