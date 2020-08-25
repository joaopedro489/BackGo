package main

import (
  "github.com/gin-gonic/gin"
  "github.com/gin-contrib/cors"
  DB "github.com/joaopedro489/BackGo/Database"
  controller "github.com/joaopedro489/BackGo/Controllers"
)

func main() {
	router := gin.Default()
  router.Use(cors.Default())
  DB.ConnectDatabase()

	router.GET("/getAllUsers", controller.getAllUsers)
  router.GET("/getUser/:id", controller.getUser)
  router.POST("/createUser", controller.createUser)
  router.PUT("/updateUser/:id", controller.UpdateUser)
  route.DELETE("/deleteUser/:id", controller.deleteUser)

  router.GET("/getAllPosts", controller.getAllPosts)
  router.GET("/getPost/:id", controller.getPost)
  router.POST("/createPost", controller.createPost)
  router.PUT("/updatePost/:id", controller.updatePost)
  route.DELETE("/deletePost/:id", controller.deletePost)


	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
