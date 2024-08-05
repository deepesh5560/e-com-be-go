package main

import (
	conterollers "e-com-be-go/controllers"
	initializers "e-com-be-go/initilizers"
	"e-com-be-go/middlewares"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}
func main() {
	r := gin.Default()
	r.POST("/user/signup", conterollers.Signup)
	r.GET("/user/login", conterollers.Login)
	r.POST("/product/add", middlewares.ValidateToken, conterollers.AddProduct)
	// r.GET("/user/getAll", conterollers.FindAllUsers)
	// r.PUT("/user/update/:id", conterollers.UpdateUser)
	// r.DELETE("/user/delete/:id", conterollers.DeleteUser)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
