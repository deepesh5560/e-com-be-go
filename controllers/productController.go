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
	role, exist := c.Get("role")
	if !exist || role != "admin" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

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

func GetAllProducts(c *gin.Context) {

	var Product []models.Product
	result := initializers.DB.Find(&Product)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": Product,
	})
}

func GetProductDtails(c *gin.Context) {
	role, exist := c.Get("role")
	if !exist || role != "admin" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	id := c.Param("id")
	var product models.Product
	result := initializers.DB.Where("id =?", id).First(&product)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": product,
	})
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	result := initializers.DB.Where("id =?", id).Delete(&product)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": "Product deleted successfully"})
}
