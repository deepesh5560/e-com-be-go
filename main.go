package main

import (
	conterollers "e-com-be-go/controllers"
	initializers "e-com-be-go/initilizers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}
func main() {
	r := gin.Default()
	r.POST("/user/create", conterollers.CreateUser)
	r.GET("/user/getAll", conterollers.FindAllUsers)
	r.PUT("/user/update/:id", conterollers.UpdateUser)
	r.DELETE("/user/delete/:id", conterollers.DeleteUser)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
