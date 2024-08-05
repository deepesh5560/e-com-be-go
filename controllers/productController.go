package conterollers

import (
	initializers "e-com-be-go/initilizers"
	"e-com-be-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddProduct(c *gin.Context) {
	var body struct {
		Name  string
		Price int
		Stock int
	}
	c.Bind(&body)
	product := models.Product{Name: body.Name, Price: body.Price, Stock: body.Stock}
	result := initializers.DB.Create(&product)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": product,
	})
}
