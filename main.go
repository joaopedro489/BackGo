package main

import (
  "github.com/gin-gonic/gin"
  "github.com/gin-contrib/cors"
  md "BackGo/Models"
  controller "BackGo/Controllers"
)

func main() {
	router := gin.Default()
  router.Use(cors.Default())
  md.ConnectDatabase()

  //CRUD de User + Funções de Relacionamento do User com Post
	router.GET("/getAllUsers", controller.GetAllUsers)
  router.GET("/getUser/:id", controller.GetUser)
  router.GET("/getUserPosts/:id", controller.GetUserPosts)
  router.POST("/userPost/:id", controller.UserPost)
  router.POST("/createUser", controller.CreateUser)
  router.PUT("/updateUser/:id", controller.UpdateUser)
  router.PUT("/userUpdatePost/:post_id/:user_id", controller.UserUpdatePost)
  router.DELETE("/deleteUser/:id", controller.DeleteUser)
  router.DELETE("userDeletePost/:post_id/:user_id", controller.UserDeletePost)

  //CRUD de Post
  router.GET("/getAllPosts", controller.GetAllPosts)
  router.GET("/getPost/:id", controller.GetPost)
  router.POST("/createPost", controller.CreatePost)
  router.PUT("/updatePost/:id", controller.UpdatePost)
  router.DELETE("/deletePost/:id", controller.DeletePost)


	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
