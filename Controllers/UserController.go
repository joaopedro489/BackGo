package Controllers

import(
  "github.com/gin-gonic/gin"
  md "BackGo/Models"
  "net/http"

)
//struct necessário para pegar o input no PUT
type UserInput struct{
   Name string `form:"name"`
   Age int `form:"age"`
   Email string `form:"email"`
   Password string `form:"-" `
}
func GetAllUsers(c *gin.Context){
   var users []md.User
   md.DB.Find(&users)
   c.JSON(http.StatusOK, gin.H{"data" : users})
}
func GetUser(c *gin.Context){
  var user md.User
  user.GetUser(c)
  c.JSON(http.StatusOK, gin.H{"data" : user})
}
func CreateUser(c *gin.Context){
    var user md.User
    c.JSON(http.StatusOK, gin.H{"data" : user.CreateUser(c)})
}
func UpdateUser(c *gin.Context){
  var user md.User
  user.GetUser(c)
  // Validate input
  var input UserInput
  if err := c.ShouldBindQuery(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }
  user.UpdateUser(input)
  c.JSON(http.StatusOK, gin.H{"data": user})

}
func DeleteUser(c *gin.Context){
    var user md.User
    user.GetUser(c)
    user.DeleteUser()
    c.JSON(http.StatusOK, gin.H{"data": "Deletado"})
}

func UserPost(c *gin.Context){
    var user md.User
    var post md.Post
    user.GetUser(c)
    var user_id uint = uint(user.ID)
    c.JSON(http.StatusOK, gin.H{"user": user,"post": post.CreatePost(c, user_id)})
}
func GetUserPosts(c *gin.Context){
  var user md.User
  var posts []md.Post
  user.GetUser(c)
  user.GetPosts(&posts)
  c.JSON(http.StatusOK, gin.H{"user": user,"posts": posts})
}
func UserDeletePost(c *gin.Context){
  var post md.Post
  var user md.User
  post.GetPost(c)
  user.GetUser(c)
  post.DeletePost()
  var posts []md.Post
  user.GetPosts(&posts)
  c.JSON(http.StatusOK, gin.H{"user": user,"posts": posts})
}
func UserUpdatePost(c *gin.Context){
  var post md.Post
  var postInput PostInput
  var user md.User
  post.GetPost(c)
  user.GetUser(c)
  if err := c.ShouldBindQuery(&postInput); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }
  post.UpdatePost(postInput)
  c.JSON(http.StatusOK, gin.H{"user": user,"post": post})

}
