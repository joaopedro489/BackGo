package Controllers

import(
  "github.com/gin-gonic/gin"
  md "BackGo/Models"
  "net/http"
  "strconv"

)
type UserInput struct{
   Name string `json:"name"`
   Age int `json:"age"`
   Email string `json:"email"`
   Password string `json:"-" `
}
func GetAllUsers(c *gin.Context){
   var users []md.User
   md.DB.Find(&users)
   c.JSON(http.StatusOK, gin.H{"data" : users})
}

func GetUser(c *gin.Context){
  var user md.User
  if err := md.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil{
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }
  var posts []md.Post
  ss := md.DB.Model(&user).Related(&posts)
  c.JSON(http.StatusOK, gin.H{"data" : ss})
}

func CreateUser(c *gin.Context){
    var user md.User
    age,_ := strconv.Atoi(c.PostForm("age"));
    user = md.User{Name: c.PostForm("name"), Age: age, Email: c.PostForm("email"), Password: c.PostForm("password")}
    md.DB.Create(&user)
    md.DB.Save(&user)
    c.JSON(http.StatusOK, gin.H{"data" : user})
}

func UpdateUser(c *gin.Context){
  var user md.User
  if err := md.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  // Validate input
  var input UserInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }
  md.DB.Model(&user).Updates(input)
  c.JSON(http.StatusOK, gin.H{"data": user})

}

func DeleteUser(c *gin.Context){
    var user md.User
    if err := md.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
       c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
       return
    }
    md.DB.Unscoped().Delete(&user)
    c.JSON(http.StatusOK, gin.H{"data": "Deletado"})
}

func UserPost(c *gin.Context){
    var user md.User
    var post md.Post
    if err := md.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
       c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
       return
    }
    post = md.Post{Tilte: c.PostForm("title"), Text: c.PostForm("text"), User_id: user.ID}
    md.DB.Create(&post)
    md.DB.Save(&post)
    c.JSON(http.StatusOK, gin.H{"user": user,"post": post})
}
func GetUserPosts(c *gin.Context){
  var user md.User
  var posts []md.Post
  if err := md.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
     c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
     return
  }
  md.DB.Where("user_id=?", user.ID).Find(&posts)
  c.JSON(http.StatusOK, gin.H{"user": user,"posts": posts})
}
func UserDeletePost(c *gin.Context){
  var post md.Post
  if err := md.DB.Where("id = ?", c.Param("post_id")).First(&post).Error; err != nil {
     c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
     return
  }
  md.DB.Delete(&post)
  c.JSON(http.StatusOK, gin.H{"data": "Deletado"})
}
