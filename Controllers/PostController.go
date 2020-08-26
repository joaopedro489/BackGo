package Controllers

import(
  "github.com/gin-gonic/gin"
  md "BackGo/Models"
  "net/http"
  "strconv"
)
//struct necessário para pegar o input no PUT
type PostInput struct{
   Title string `form:"title"`
   Text string `form:"text"`
}

func GetAllPosts(c *gin.Context){
   var posts []md.Post
   md.DB.Find(&posts)
   c.JSON(http.StatusOK, gin.H{"data" : posts})
}

func GetPost(c *gin.Context){
  var post md.Post
  post.GetPost(c)
  c.JSON(http.StatusOK, gin.H{"data" : post})
}

func CreatePost(c *gin.Context){
    var post md.Post
    user_id,_ := strconv.Atoi(c.PostForm("user_id"));
    userId := uint(user_id)
    c.JSON(http.StatusOK, gin.H{"data": post.CreatePost(c, userId)})
}

func UpdatePost(c *gin.Context){
  var post md.Post
  post.GetPost(c)
  // Validate input
  var input PostInput
  if err := c.ShouldBindQuery(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }
  post.UpdatePost(input)
  c.JSON(http.StatusOK, gin.H{"data": post})
}

func DeletePost(c *gin.Context){
    var post md.Post
    post.GetPost(c)
    post.DeletePost()
    c.JSON(http.StatusOK, gin.H{"data": "Deletado"})
}
