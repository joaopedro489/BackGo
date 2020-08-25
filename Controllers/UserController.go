package Controllers

import(
  "github.com/gin-gonic/gin"
  DB "github.com/joaopedro489/BackGo/Database"
  md "github.com/joaopedro489/BackGo/Models"
)
func getAllUsers(c *gin.Context){
   var users []md.User
   DB.DB.Find(&users)
   c.JSON(http.StatusOK, gin.H{"data" : users})
}

func getUser(c *gin.Context){
  var user []md.User
  if err := DB.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil{
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }
  c.JSON(http.StatusOK, gin.H{"data" : user})
}

func createUser(c *gin.Context){
    var userInput md.User
    if err := c.ShouldBindJSON(&userInput); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    var user md.User
    user.createUser(userInput)
    c.JSON(http.StatusOK, gin.H{"data": user})
}
func updateUser(c *gin.Context){
    var user md.User
    if err := DB.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }
    var userInput md.User
    if err := c.ShouldBindJSON(&userInput); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    user.UpdateUser(userInput)
    c.JSON(http.StatusOK, gin.H{"data" : user})
}
func deleteUser(c *gin.Context){
    var user md.User
    if err := DB.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
       c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
       return
    }
    user.deleteUser();
    c.JSON(http.StatusOK, gin.H{"data": "Deletado"})
}
