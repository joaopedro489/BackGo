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

	router.GET("/getAllUsers", controller.GetAllUsers)
  router.GET("/getUser/:id", controller.GetUser)
  router.GET("/getUserPosts/:id", controller.GetUserPosts)
  router.POST("/UserPost/:id", controller.UserPost)
  router.POST("/createUser", controller.CreateUser)
  router.PUT("/updateUser/:id", controller.UpdateUser)
  router.DELETE("/deleteUser/:id", controller.DeleteUser)
  router.DELETE("UserDeletePost/:post_id", controller.UserDeletePost)

  router.GET("/getAllPosts", controller.GetAllPosts)
  router.GET("/getPost/:id", controller.GetPost)
  router.POST("/createPost", controller.CreatePost)
  router.PUT("/updatePost/:id", controller.UpdatePost)
  router.DELETE("/deletePost/:id", controller.DeletePost)


	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
