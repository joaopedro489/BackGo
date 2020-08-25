package Database
import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  md "github.com/joaopedro489/BackGo/Models"
)
var DB *gorm.DB

func ConnectDatabase(){
  db,err := gorm.Open("mysql", "root:@/demo?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
  panic("failed to connect database")
 }
 db.AutoMigrate(&md.Post{},&md.User{})
 DB = db
}
