package Controllers

import(
  "github.com/gin-gonic/gin"
  md "BackGo/Models"
  "net/http"
  "strconv"
)
type PostInput struct{
   Title string `json:"title"`
   Text string `json:"text"`
}
func GetAllPosts(c *gin.Context){
   var posts []md.Post
   md.DB.Find(&posts)
   c.JSON(http.StatusOK, gin.H{"data" : posts})
}

func GetPost(c *gin.Context){
  var post md.Post
  if err := md.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil{
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }
  c.JSON(http.StatusOK, gin.H{"data" : post})
}

func CreatePost(c *gin.Context){
    var post md.Post
    user_id,_ := strconv.Atoi(c.PostForm("user_id"));
    userId := uint(user_id)
    post = md.Post{Tilte: c.PostForm("title"), Text: c.PostForm("text"), User_id: userId}
    md.DB.Create(&post)
    md.DB.Save(&post)
    c.JSON(http.StatusOK, gin.H{"data": post})
}
func UpdatePost(c *gin.Context){
  var post md.Post
  if err := md.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  // Validate input
  var input PostInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }
  md.DB.Model(&post).Updates(input)
  c.JSON(http.StatusOK, gin.H{"data": post})

}

func DeletePost(c *gin.Context){
    var post md.Post
    if err := md.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
       c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
       return
    }
    md.DB.Unscoped().Delete(&post)
    c.JSON(http.StatusOK, gin.H{"data": "Deletado"})
}
