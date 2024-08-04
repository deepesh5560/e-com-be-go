package conterollers

import (
	initializers "e-com-be-go/initilizers"
	"e-com-be-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var body struct {
		Name  string
		Email string
		Age   int
	}
	c.Bind(&body)
	user := models.User{Name: body.Name, Email: body.Email, Age: body.Age}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": user,
	})
}

func FindAllUsers(c *gin.Context) {

	var users []models.User
	result := initializers.DB.Find(&users)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": users,
	})
}

func UpdateUser(c *gin.Context) {
	var body struct {
		Name  string
		Email string
		Age   int
	}
	c.Bind(&body)

	id := c.Params.ByName("id")
	var user models.User
	initializers.DB.First(&user, id)

	result := initializers.DB.Model(&user).Updates(models.User{Name: body.Name, Email: body.Email, Age: body.Age})

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": user,
	})
}

func DeleteUser(c *gin.Context) {

	id := c.Params.ByName("id")

	initializers.DB.Delete(&models.User{}, id)

	c.JSON(http.StatusOK, gin.H{
		"result": "user deleted",
	})
}
