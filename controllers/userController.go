package conterollers

import (
	initializers "e-com-be-go/initilizers"
	"e-com-be-go/middlewares"
	"e-com-be-go/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// func CreateUser(c *gin.Context) {
// 	var body struct {
// 		Name  string
// 		Email string
// 		Age   int
// 	}
// 	c.Bind(&body)
// 	user := models.User{Name: body.Name, Email: body.Email, Age: body.Age}
// 	result := initializers.DB.Create(&user)

// 	if result.Error != nil {
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"result": user,
// 	})
// }

// func FindAllUsers(c *gin.Context) {

// 	var users []models.User
// 	result := initializers.DB.Find(&users)

// 	if result.Error != nil {
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"result": users,
// 	})
// }

// func UpdateUser(c *gin.Context) {
// 	var body struct {
// 		Name  string
// 		Email string
// 		Age   int
// 	}
// 	c.Bind(&body)

// 	id := c.Params.ByName("id")
// 	var user models.User
// 	initializers.DB.First(&user, id)

// 	result := initializers.DB.Model(&user).Updates(models.User{Name: body.Name, Email: body.Email, Age: body.Age})

// 	if result.Error != nil {
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"result": user,
// 	})
// }

// func DeleteUser(c *gin.Context) {

// 	id := c.Params.ByName("id")

// 	initializers.DB.Delete(&models.User{}, id)

// 	c.JSON(http.StatusOK, gin.H{
// 		"result": "user deleted",
// 	})
// }

func Signup(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}
	c.Bind(&body)
	hashPass, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		panic(err)
	}
	user := models.User{Email: body.Email, Password: string(hashPass)}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	token, err := middlewares.CreateToken(body.Email)
	fmt.Println("token", token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Unable to create token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": user,
		"token":  token,
	})
}

func Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}
	c.Bind(&body)

	var user models.User
	initializers.DB.First(&user, "Email = ?", body.Email)

	if user.ID == 0 {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	// result := initializers.DB.Create(&user)
	if err != nil {

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Wrong username or password"})
		return

	}

	token, err := middlewares.CreateToken(body.Email)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Unable to create token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": user,
		"token":  token,
	})
}
