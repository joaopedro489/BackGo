package Controllers

import(
  "github.com/gin-gonic/gin"
  DB "github.com/joaopedro489/BackGo/Database"
  md "github.com/joaopedro489/BackGo/Models"
)

func getAllPosts(c *gin.Context){
   var posts []md.Post
   DB.DB.Find(&posts)
   c.JSON(http.StatusOK, gin.H{"data" : posts})
}

func getPost(c *gin.Context){
  var post []md.Post
  if err := DB.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil{
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }
  c.JSON(http.StatusOK, gin.H{"data" : post})
}

func createPost(c *gin.Context){
    var postInput md.Post
    if err := c.ShouldBindJSON(&postInput); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    var post md.Post
    post.createPost(postInput)
    c.JSON(http.StatusOK, gin.H{"data": post})
}
func updatePost(c *gin.Context){
    var post md.Post
    if err := DB.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }
    var postInput md.Post
    if err := c.ShouldBindJSON(&postInput); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    post.UpdatePost(postInput)
    c.JSON(http.StatusOK, gin.H{"data" : post})
}
func deletePost(c *gin.Context){
    var post md.Post
    if err := DB.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
       c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
       return
    }
    post.deletePost();
    c.JSON(http.StatusOK, gin.H{"data": "Deletado"})
}
